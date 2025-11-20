package handlers

import (
	"net/http"
	"strconv"

	"API/data"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, data.GetItems())
}

func GetItemByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if item, ok := data.GetItemByID(id); ok {
		c.JSON(http.StatusOK, item)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}
