package server

import (
	"plantain/server/controller/monitor"
	"plantain/server/controller/rtdb"

	"github.com/gin-gonic/gin"
)

var RestartChan = make(chan interface{}, 1)

func RouterWeb(port string) {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, This is Plantain Configuration Tools.")
	})

	// RTDB相关配置
	rtdbRouter := router.Group("/api/v1/rtdb")
	{
		// System Collector
		rtdbRouter.POST("/restart", func(c *gin.Context) {
			RestartChan <- 1
		})

		// CURD Collector
		rtdbRouter.GET("/collector/list", rtdb.ApiGetCollectorList)
		rtdbRouter.GET("/collector/id/:id", rtdb.ApiGetCollectorById)
		rtdbRouter.GET("/collectorWithRtTableSet/id/:id",
			rtdb.ApiGetCollectorByIdWithRtTableSet)
		rtdbRouter.GET("/collectorWithRtTableSet/name/:name",
			rtdb.ApiGetCollectorByNameWithRtTableSet)
		rtdbRouter.GET("/collector/name/:name", rtdb.ApiGetCollectorByName)
		rtdbRouter.POST("/collector", rtdb.ApiAddCollectorItemInList)
		rtdbRouter.POST("/collectorAndRtTable", rtdb.ApiAddCollectorItemInListAndCreateRtTable)
		rtdbRouter.DELETE("/collector/name/:name", rtdb.ApiDelCollectorItemInListByName)
		rtdbRouter.DELETE("/collectorAndRtTable/name/:name", rtdb.ApiDelCollectorItemInListAndDropRtTableByName)
		rtdbRouter.DELETE("/collector/id/:id", rtdb.ApiDelCollectorItemInListById)
		rtdbRouter.DELETE("/collectorAndRtTable/id/:id", rtdb.ApiDelCollectorItemInListAndDropRtTableById)
		rtdbRouter.PUT("/collector/id/:id", rtdb.ApiUpdateCollectorItemInListById)
		rtdbRouter.PUT("/collector/name/:name", rtdb.ApiUpdateCollectorItemInListByName)

		// CURD RtTable
		rtdbRouter.POST("/rtTable/create/:tableName", rtdb.ApiCreateRTTable)
		rtdbRouter.DELETE("/rtTable/drop/:tableName", rtdb.ApiDropRTTable)
		rtdbRouter.GET("/rtTable/:tableName", rtdb.ApiGetRTTable)
		rtdbRouter.POST("/rtTable/item/:tableName", rtdb.ApiAddItemInRTTable)
		rtdbRouter.DELETE("/rtTable/item/:tableName/:pid", rtdb.ApiDelItemInRTTableByPID)
		rtdbRouter.PUT("/rtTable/item/:tableName/:pid", rtdb.ApiUpdateItemInRTTable)
	}

	// 查看实时值、报警、历史值
	monitorRouter := router.Group("/api/v1/monitor")
	{
		monitorRouter.GET("rtValue/byRtTable/:tableName", monitor.ApiGetRtTableRealTimeValueByTableName)
		monitorRouter.GET("rtValue/byPID/:pid", monitor.ApiGetRealTimeValueByPID)
	}
	router.Run(port)
}
