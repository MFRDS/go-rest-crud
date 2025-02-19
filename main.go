package main

import (
	"log"

	"go-rest-crud/controllers"
	"go-rest-crud/database"
	"go-rest-crud/repository"
	"go-rest-crud/routes"
	"go-rest-crud/services"
)

func main() {

	database.ConnectDatabase()

	repo := repository.NewMahasiswaRepository()
	service := services.NewMahasiswaService(repo)
	controller := controllers.NewMahasiswaController(service)

	r := routes.SetupRouter(controller)

	log.Println("Server running on port 8080")
	r.Run(":8080")
}
