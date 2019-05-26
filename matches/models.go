package matches

import (
    "github.com/jinzhu/gorm"

    "github.com/resinderate/ping-pong-backend/common"
)

type MatchModel struct {
    gorm.Model
    Winner string `json:"winner"`
    Loser string `json:"loser"`
    WinnerGain uint `json:"winner_gain"`
    LoserLoss uint `json:"loser_loss"`
}

func AutoMigrate() {
    db := common.GetDB()
    db.AutoMigrate(&MatchModel{})
}
