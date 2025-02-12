package main

import (
	"go-rest-crud/database"
	"go-rest-crud/routes"
)

func main() {
	database.ConnectDatabase()
	r := routes.SetupRouter()
	r.Run("0.0.0.0:8080")

}
