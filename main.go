/**
 * Created by cheenwe on 20200528.
 */
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"10sh.cn/ip/ip17mon"
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
func CheckShorter(c *gin.Context) {
	// c.JSON(200, devices)

	pth := fmt.Sprintf("%s", c.Request.URL)
	fmt.Println(strings.Count(pth, "/")) //2
	//截取
	if strings.Count(pth, "/") == 1 {
		c.String(http.StatusOK, pth[1:])
	} else {
		c.String(http.StatusNotFound, "not find")
	}
}

func main() {
	router := gin.Default()
	gin.ForceConsoleColor()

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

	router.NoRoute((func(c *gin.Context) { CheckShorter(c) }))

	router.Static("/assets", "./assets")

	router.StaticFile("/favicon.ico", "./assets/favicon.ico")

	router.Run(":8081")

}
