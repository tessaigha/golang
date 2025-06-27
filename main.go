package main

import (
    "go-kasir-api/routes"
    "go-kasir-api/database"
)

func main() {
    database.InitDB()
    r := routes.SetupRouter()
    r.Run(":8080")
}
