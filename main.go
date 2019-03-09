package main

import (
    // "fmt"

    "github.com/gin-gonic/gin"
    // "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"

    "github.com/resinderate/ping-pong-backend/common"
    "github.com/resinderate/ping-pong-backend/players"
    "github.com/resinderate/ping-pong-backend/matches"
)

func main() {
    db := common.InitDB()

    // Batch into migrate step.
    // Perhaps let the apps manage their own migrations.
    db.AutoMigrate(&players.PlayerModel{})
    db.AutoMigrate(&matches.MatchModel{})

	routes := gin.Default()
    routes.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello World"})
    })
    routes.Run()
}
