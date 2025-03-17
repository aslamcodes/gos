package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/aslamcodes/gos/database"
)


func main() {
    database.InitDB()
    
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello World Again")
    })
    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
