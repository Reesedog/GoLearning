package database

import (
    "log"
    "memo/models"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func InitDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("memo.db"), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    db.AutoMigrate(&models.Memo{})

    return db
}
