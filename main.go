package main

import (
    "github.com/gin-gonic/gin"
    "API/handlers"
)

func main() {
    r := gin.Default()

    r.GET("/items", handlers.GetItems)
    r.GET("/items/:id", handlers.GetItemByID)

    r.Run(":8080")
}
