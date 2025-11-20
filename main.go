package main

import (
	"API/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/items", handlers.GetItems)
	r.GET("/items/:id", handlers.GetItemByID)

	r.Run(":8080")
}
