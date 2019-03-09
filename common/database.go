package common

import (
    "fmt"
    "github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
    db, err := gorm.Open("sqlite3", "./ping-pong.db")
    if err != nil {
        fmt.Println("db err:", err)
    }
    db.DB().SetMaxIdleConns(10)
    DB = db
    return DB
}

func GetDB() *gorm.DB {
    return DB
}
