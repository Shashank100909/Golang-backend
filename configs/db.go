package configs

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "host=localhost user=postgres password=rudransh dbname=Go_db port=5432 sslmode=disable TimeZone=Asia/Kolkata"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }
    DB = db
    fmt.Println("Database connected!")
}