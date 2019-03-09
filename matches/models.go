package matches

import (
    "github.com/jinzhu/gorm"
)

type MatchModel struct {
    gorm.Model
    Winner string `json:"winner"`
    Loser string `json:"loser"`
    WinnerGain uint `json:"winner_gain"`
    LoserLoss uint `json:"loser_loss"`
}
