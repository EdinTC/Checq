package main

import (
	"github.com/EdinTC/Checq/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/lookup-ip/:name", utils.QueryIPAdress)
	router.GET("/lookup-txt/:name", utils.QueryTXT)
	router.GET("/lookup-ns/:name", utils.QueryNS)
	router.Run(":3000")
}
