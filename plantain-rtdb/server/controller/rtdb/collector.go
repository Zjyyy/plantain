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

func ApiGetCollectorByIdWithRtTableSet(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	data, err := models.GetCollectorByIdWithRtTableSet(id)

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
	data, err := models.GetCollectorByName(name)

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
func ApiGetCollectorByNameWithRtTableSet(c *gin.Context) {
	name := c.Param("name")
	data, err := models.GetCollectorByNameWithRtTableSet(name)

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

func ApiAddCollectorItemInListAndCreateRtTable(c *gin.Context) {
	var data request.CollectorReq
	c.BindJSON(&data)

	err := models.AddCollectorItemInListAndCreateRtTable(&models.Collector{
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

func ApiDelCollectorItemInListAndDropRtTableById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := models.DelCollectorItemInListAndDropRtTableById(id)

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

func ApiDelCollectorItemInListAndDropRtTableByName(c *gin.Context) {
	name := c.Param("name")

	err := models.DelCollectorItemInListAndDropRtTableByName(name)

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

func ApiUpdateCollectorItemInListByName(c *gin.Context) {
	name := c.Param("name")

	var data request.CollectorReq
	c.BindJSON(&data)

	err := models.UpdateCollectorItemInListByName(name, &models.Collector{
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
		"message": fmt.Sprintf("%v", err),
		"data":    true,
	})
}

func ApiUpdateCollectorItemInListById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	var data request.CollectorReq
	c.BindJSON(&data)

	err := models.UpdateCollectorItemInListById(id, &models.Collector{
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
		"message": fmt.Sprintf("%v", err),
		"data":    true,
	})
}
