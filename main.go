package main

import (
	"os"

	"github.com/wujie1993/sonarqube-dingtalk/dingtalk"

	"github.com/gin-gonic/gin"
)

var (
	DingTalkAccessToken string
	ListenAddr          string
)

func init() {
	DingTalkAccessToken = os.Getenv("dingtalk_access_token")
	ListenAddr = os.Getenv("listen_addr")
	if ListenAddr == "" {
		ListenAddr = "localhost:8080"
	}
}

func main() {

	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		//projectKey:=c.GetQuery("id")
		if err := dingtalk.SendJSON(DingTalkAccessToken, ""); err != nil {
			c.AbortWithError(500, err)
			return
		}
		c.JSON(200, gin.H{})
	})
	r.Run(ListenAddr)
}
