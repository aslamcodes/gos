package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "os"
    "github.com/aslamcodes/gos/internal/donor"
)


func main() {
    db_host := os.Getenv("DB_HOST")
    db_port := os.Getenv("DB_PORT")
    db_user := os.Getenv("DB_USER")
    db_pass := os.Getenv("DB_PASSWORD")
    db_name := os.Getenv("DB_NAME")
    
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_host, db_user, db_pass, db_name, db_port)


    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("Failed to connect to database")
    }

    fmt.Println("Connection established with Database") 

    db.AutoMigrate(&donor.Donor{})
    
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello World Again")
    })
    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
