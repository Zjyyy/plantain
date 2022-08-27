package rtdb

import (
	"fmt"
	"net/http"
	"plantain/models"
	"plantain/models/dtos/request"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ApiGetCollectorList(c *gin.Context) {
	data, err := models.GetAllCollectorList()

	c.JSON(http.StatusBadRequest, gin.H{
		"status":  true,
		"message": fmt.Sprintf("%v", err),
		"data":    data,
	})
}

func ApiGetCollectorById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	data, err := models.GetCollectorById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": fmt.Sprintf("%v", err),
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "",
		"data":    data,
	})
}

func ApiGetCollectorByName(c *gin.Context) {
	name := c.Param("name")
	data, err := models.GetCollectorByCollectorName(name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": fmt.Sprintf("%v", err),
			"data":    "",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  true,
		"message": fmt.Sprintf("%v", err),
		"data":    data,
	})
}

func ApiAddCollectorItemInList(c *gin.Context) {
	var data request.CollectorReq
	c.BindJSON(&data)

	err := models.AddCollectorItemInList(&models.Collector{
		CollectorName: data.CollectorName,
		Version:       data.Version,
		DllPath:       data.DllPath,
		ConnStr:       data.ConnStr,
		Setting:       data.Setting,
		Des:           data.Des,
		RtTableName:   data.RtTableName,
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
		"message": "",
		"data":    true,
	})
}

func ApiDelCollectorItemInListById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := models.DelCollectorItemInListById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": fmt.Sprintf("%v", err),
			"data":    false,
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status":  true,
		"message": fmt.Sprintf("%v", err),
		"data":    true,
	})
}

func ApiDelCollectorItemInListByName(c *gin.Context) {
	name := c.Param("name")

	err := models.DelCollectorItemInListByName(name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": fmt.Sprintf("%v", err),
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status":  true,
		"message": fmt.Sprintf("%v", err),
		"data":    "",
	})
}
