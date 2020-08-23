package main

import (
	"github.com/EdinTC/Checq/api/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	router.Use(cors.New(config))
	router.GET("/:name", handlers.QueryHandler)
	// router.GET("/lookup-ip/:name", utils.QueryIPAddress)
	// router.GET("/lookup-txt/:name", utils.QueryTXT)
	// router.GET("/lookup-ns/:name", utils.QueryNS)
	router.Run(":1337")
}
