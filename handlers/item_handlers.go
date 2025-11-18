package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "API/data"
)

func GetItems(c *gin.Context) {
    c.JSON(http.StatusOK, data.Items)
}

func GetItemByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    for _, item := range data.Items {
        if item.ID == id {
            c.JSON(http.StatusOK, item)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}
