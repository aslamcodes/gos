package database

import (
	"gorm.io/gorm"
    "os"
    "fmt"
    "github.com/aslamcodes/gos/internal/donor" 
    "gorm.io/driver/postgres"
)

var DB *gorm.DB

func InitDB() {
    db_host := os.Getenv("DB_HOST")
    db_port := os.Getenv("DB_PORT")
    db_user := os.Getenv("DB_USER")
    db_pass := os.Getenv("DB_PASSWORD")
    db_name := os.Getenv("DB_NAME")
    
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_host, db_user, db_pass, db_name, db_port)

    DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("Failed to connect to database")
    }

    fmt.Println("Connection established with Database") 

    DB.AutoMigrate(&donor.Donor{})
}

