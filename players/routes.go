package players

import (
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
    router.POST("/", CreatePlayer)
    router.GET("/", ListPlayers)
    router.GET("/:id/", GetPlayer)
}

func CreatePlayer(c *gin.Context) {
    c.JSON(200, gin.H{"message": "Player Created!"})
}

func ListPlayers(c *gin.Context) {
    players := []interface{}{"Player 1", "Player 2"}
    c.JSON(200, players)
}

func GetPlayer(c *gin.Context) {
    id := c.Param("id")
    c.JSON(200, gin.H{"message": "Single Player -- " + id})
}
