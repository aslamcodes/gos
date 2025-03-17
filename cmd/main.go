package main

import (
    "gos/config"
    "gos/routes"
)


func main() {
    config.InitDB()
    
    r := routes.SetupRouter()

    r.Run(":8080")
}
