package server

import (
	"plantain/server/contorl"

	"github.com/gin-gonic/gin"
)

func RouterWeb(port string) {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, This is Plantain Configuration Tools.")
	})

	driverRouter := router.Group("/api/v1/rtdb")
	{
		driverRouter.GET("/driverList", contorl.ApiReadAllDriver)
	}

	router.Run(port)
}
