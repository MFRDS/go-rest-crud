package main

import (
	"github.com/MFRDS/go-rest-crud/database"
	"github.com/MFRDS/go-rest-crud/routes"
)

func main() {
	database.ConnectDatabase()
	r := routes.SetupRouter()
	r.Run("0.0.0.0:8080")

}
