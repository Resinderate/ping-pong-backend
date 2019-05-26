package players

import (
    "github.com/jinzhu/gorm"

    "github.com/resinderate/ping-pong-backend/common"
)

type PlayerModel struct {
    gorm.Model
    Name string `json:"name"`
    Elo uint `json:"elo"`
}

func AutoMigrate() {
    db := common.GetDB()
    db.AutoMigrate(&PlayerModel{})
}