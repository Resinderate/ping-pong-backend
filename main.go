package main

import (
    "github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/sqlite"

    "github.com/resinderate/ping-pong-backend/common"
    "github.com/resinderate/ping-pong-backend/players"
    "github.com/resinderate/ping-pong-backend/matches"
)

func main() {
    common.InitDB()
    players.AutoMigrate()
    matches.AutoMigrate()

	routes := gin.Default()
    AddRoutes(routes)

    routes.Run()
}

func AddRoutes(routes *gin.Engine) {
    playersGroup := routes.Group("/players")
    players.RegisterRoutes(playersGroup)

    matchesGroup := routes.Group("/matches")
    matches.RegisterRoutes(matchesGroup)
}