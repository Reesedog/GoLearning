package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    database, err := gorm.Open(sqlite.Open("memo.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&models.Memo{})

    DB = database
}
