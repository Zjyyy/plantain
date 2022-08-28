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
		rtdbRouter.PUT("/collector/id/:id", rtdb.ApiUpdateCollectorItemInListById)
		rtdbRouter.PUT("/collector/name/:name", rtdb.ApiUpdateCollectorItemInListByName)

		rtdbRouter.POST("/rtTable/create/:tableName", rtdb.ApiCreateRTTable)
		rtdbRouter.DELETE("/rtTable/drop/:tableName", rtdb.ApiDropRTTable)
		rtdbRouter.GET("/rtTable/:tableName", rtdb.ApiGetRTTable)
		rtdbRouter.POST("/rtTable/item/:tableName", rtdb.ApiAddItemInRTTable)
		rtdbRouter.DELETE("/rtTable/item/:tableName/:pid", rtdb.ApiDelItemInRTTableByPID)
		rtdbRouter.PUT("/rtTable/item/:tableName/:pid", rtdb.ApiUpdateItemInRTTable)
	}

	router.Run(port)
}
