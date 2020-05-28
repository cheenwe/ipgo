/**
 * Created by cheenwe on 20200528.
 */
package main

import (
	"net"
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

func main() {
	router := gin.Default()

	router.GET("/ip", func(c *gin.Context) {
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
	})

	router.Static("/assets", "./assets")

	router.StaticFile("/favicon.ico", "./assets/favicon.ico")

	router.Run(":8081")

}
