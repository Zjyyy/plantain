package monitor

import (
	"fmt"
	"net/http"
	"plantain/core"

	"github.com/gin-gonic/gin"
)

func ApiGetRtTableRealTimeValueList(c *gin.Context) {
	collectorName := c.Param("collectorName")
	handler := core.MemoryBlockHandler
	if handler == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "MemoryBlockHandler Error.It is Nil.",
			"data":    "",
		})
		return
	}
	cache, err := handler.GetCache(collectorName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": fmt.Sprintf("GetCache Error:%v", err),
			"data":    "",
		})
		return
	}
	type temp struct {
		Key   string
		Value interface{}
	}
	items := cache.Items()
	result := make([]temp, len(items))
	for key, val := range cache.Items() {
		result = append(result, temp{key, val.Object})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "",
		"data":    result,
	})
}

func ApiGetRtTableRealTimeValue(c *gin.Context) {
	collectorName := c.Param("collectorName")
	pid := c.Param("pid")
	handler := core.MemoryBlockHandler
	if handler == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "MemoryBlockHandler Error.It is Nil.",
			"data":    "",
		})
		return
	}
	val, err := handler.ReadFromCache(collectorName, pid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": fmt.Sprintf("ReadFromCache Error:%v", err),
			"data":    "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "",
		"data":    val,
	})
}
