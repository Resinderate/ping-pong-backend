package players

import (
    "github.com/jinzhu/gorm"
)

type PlayerModel struct {
    gorm.Model
    Name string `json:"name"`
    Elo uint `json:"elo"`
}
