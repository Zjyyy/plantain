package rtdb

import (
	"fmt"
	"net/http"
	"plantain/models"
	"plantain/models/dtos/request"

	"github.com/gin-gonic/gin"
)

func ApiGetRTTable(c *gin.Context) {
	tableName := c.Param("tableName")

	data, err := models.GetRTTable(tableName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": fmt.Sprintf("%v", err),
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": fmt.Sprintf("%v", err),
		"data":    data,
	})
}

func ApiCreateRTTable(c *gin.Context) {
	tableName := c.Param("tableName")

	err := models.CreateRTTable(tableName)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": fmt.Sprintf("%v", err),
			"data":    false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": fmt.Sprintf("%v", err),
		"data":    true,
	})
}

func ApiAddItemInRTTable(c *gin.Context) {
	tableName := c.Param("tableName")
	var data models.RtTable
	c.ShouldBindJSON(&data)

	err := models.AddItemInRTTable(tableName, &data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": fmt.Sprintf("%v", err),
			"data":    false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "",
		"data":    true,
	})
}

func ApiDelItemInRTTableByPID(c *gin.Context) {
	tableName := c.Param("tableName")
	pid := c.Param("pid")

	err := models.DeleteItemInRTTableByPID(tableName, pid)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": fmt.Sprintf("%v", err),
			"data":    false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "",
		"data":    true,
	})
}

func ApiUpdateItemInRTTable(c *gin.Context) {
	tableName := c.Param("tableName")
	pid := c.Param("pid")

	var data request.RTTableReqDto
	c.BindJSON(&data)

	err := models.UpdateItemInRTTableByPID(tableName, pid, &models.RtTable{
		PID:          data.PID,
		Value:        data.Value,
		ValueType:    data.ValueType,
		Des:          data.Des,
		Address:      data.Address,
		LimitUp:      data.LimitUp,
		LimitDown:    data.LimitDown,
		Level:        data.Level,
		AlarmDes:     data.AlarmDes,
		IsHistorical: data.IsHistorical,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": fmt.Sprintf("%v", err),
			"data":    false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": fmt.Sprintf("%v", err),
		"data":    true,
	})
}

func ApiDropRTTable(c *gin.Context) {
	tableName := c.Param("tableName")

	err := models.DropRTTable(tableName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": fmt.Sprintf("%v", err),
			"data":    false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "",
		"data":    true,
	})
}
