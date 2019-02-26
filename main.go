package main

import "github.com/gin-gonic/gin"

func main() {
	routes := gin.Default()
    routes.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello World"})
    })
    routes.Run()
}
