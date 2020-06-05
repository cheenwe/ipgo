/**
 * Created by cheenwe on 20200528.
 */
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"10sh.cn/ip/db"
	"10sh.cn/ip/pkg/ip17mon"
	"10sh.cn/ip/pkg/shortid"
	"github.com/gin-gonic/gin"
)

// func CheckIp(ip string) bool {
// 	addr := strings.Trim(ip, " ")
// 	regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
// 	if match, _ := regexp.MatchString(regStr, addr); match {
// 		return true
// 	}
// 	return false
// }

type Res struct {
	Ip       string
	Country  string
	Province string
	City     string
}

func IpToAddress(c *gin.Context) {

	ip := c.Query("ip")
	address := net.ParseIP(ip)
	result := map[string]interface{}{}

	ret := new(Res)
	if address == nil {
		result["msg"] = " IP 格式错误"
		result["code"] = 0
		result["data"] = ret
	} else {
		info := string(ip17mon.Find(ip))
		res := strings.Split(info, "\t")
		result["msg"] = "成功"
		result["code"] = 1
		ret.Ip = ip
		ret.Country = res[0]
		ret.Province = res[1]
		ret.City = res[2]

		if ret.Country == ret.Province {
			ret.Province = ""
			ret.City = ""
		} else if ret.City == ret.Province {
			ret.City = ""
		}
		result["data"] = ret
	}
	c.JSON(200, result)
}

func SaveUrl(conn *sql.DB, c *gin.Context) {
	url := c.Query("url")
	code := shortid.MustGenerate()
	log.Printf("C: %s", url)
	log.Printf("save code: === [%s] %s", code, url)

	db.InsertLink(conn, url, code)

	result := map[string]interface{}{}
	result["msg"] = "成功"
	result["code"] = 1
	result["data"] = code
	c.JSON(200, result)
}

func CheckShorter(conn *sql.DB, c *gin.Context) {
	pth := fmt.Sprintf("%s", c.Request.URL)
	fmt.Println(strings.Count(pth, "/")) //2
	//截取
	if strings.Count(pth, "/") == 1 {
		code := pth[1:]
		log.Printf("check code: === [%s]", code)
		url := db.QueryLink(conn, code)
		if url == "0" {
			c.String(http.StatusNotFound, "Not Find")
		} else {
			c.Redirect(http.StatusMovedPermanently, url)
		}
	} else {
		c.String(http.StatusNotFound, "Not Find")
	}
}

// 1. 如果返回的错误为nil,说明文件或文件夹存在
// 2. 如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
// 3. 如果返回的错误为其它类型,则不确定是否在存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	router := gin.Default()
	gin.ForceConsoleColor()
	// 判断db目录存在，不存在则新建
	db_path := "./db/"
	db_file := db_path + "data.db"

	db_path_exist, _ := PathExists(db_path)

	if db_path_exist != true {
		os.MkdirAll(db_path, os.ModePerm)
	}
	conn := db.ConnectDB("sqlite3", db_file)

	fileInfo, _ := os.Stat(db_file)
	db_size := fileInfo.Size()
	log.Printf("db_size: %v", db_size)

	db.CreateLinkTable(conn)

	router.POST("/url", func(c *gin.Context) { SaveUrl(conn, c) })
	router.POST("/ip", func(c *gin.Context) { IpToAddress(c) })

	router.GET("/s/:code", func(c *gin.Context) {
		log.Printf("C: %v", c)
		log.Printf("code: %v", c.Param("code"))

		fmt.Println(c.Param("code"))
		fmt.Println(c)
		fmt.Println(c)
		fmt.Println(c.Query("code"))

		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	router.NoRoute((func(c *gin.Context) { CheckShorter(conn, c) }))

	router.Static("/assets", "./assets")

	router.StaticFile("/favicon.ico", "./assets/favicon.ico")

	router.Run(":8081")

}
