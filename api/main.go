package main

import (
	"github.com/EdinTC/Checq/api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/:name", handlers.QueryHandler)
	// router.GET("/lookup-ip/:name", utils.QueryIPAdress)
	// router.GET("/lookup-txt/:name", utils.QueryTXT)
	// router.GET("/lookup-ns/:name", utils.QueryNS)
	router.Run(":1337")
}
