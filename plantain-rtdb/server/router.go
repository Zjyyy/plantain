package server

import (
	"plantain/server/controller/rtdb"

	"github.com/gin-gonic/gin"
)

func RouterWeb(port string) {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, This is Plantain Configuration Tools.")
	})

	rtdbRouter := router.Group("/api/v1/rtdb")
	{
		rtdbRouter.GET("/collector/list", rtdb.ApiGetCollectorList)
		rtdbRouter.GET("/collector/id/:id", rtdb.ApiGetCollectorById)
		rtdbRouter.GET("/collector/name/:name", rtdb.ApiGetCollectorByName)
		rtdbRouter.POST("/collector", rtdb.ApiAddCollectorItemInList)
		rtdbRouter.DELETE("/collector/name/:name", rtdb.ApiDelCollectorItemInListByName)
		rtdbRouter.DELETE("/collector/id/:id", rtdb.ApiDelCollectorItemInListById)

		rtdbRouter.GET("/rtTable/:tableName", rtdb.ApiGetRTTable)
		rtdbRouter.DELETE("/rtTable/:tableName/:pid", rtdb.ApiDropRTTable)
		rtdbRouter.POST("/rtTable/item/:tableName", rtdb.ApiAddItemInRTTable)
		rtdbRouter.DELETE("/rtTable/item/:tableName", rtdb.ApiDelItemInRTTableByPID)
	}

	router.Run(port)
}
