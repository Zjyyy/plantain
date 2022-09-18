package monitor

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiGetRtTableRealTimeValueByTableName(c *gin.Context) {
	tableName := c.Param("tableName")

	c.JSON(http.StatusOK, gin.H{
		"status":  false,
		"message": fmt.Sprintf("%v", tableName),
		"data":    "",
	})
}

func ApiGetRealTimeValueByPID(c *gin.Context) {
	pid := c.Param("pid")

	c.JSON(http.StatusOK, gin.H{
		"status":  false,
		"message": fmt.Sprintf("%v", pid),
		"data":    "",
	})
}
