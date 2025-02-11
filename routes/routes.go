package routes

import (
	"github.com/MFRDS/go-rest-crud/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/mahasiswa", controller.GetMahasiswa)
	r.POST("/mahasiswa", controller.CreateMahasiswa)
	r.GET("/mahasiswa/:id", controller.GetMahasiswaByID)
	r.PUT("/mahasiswa/:id", controller.UpdateMahasiswa)
	r.DELETE("/mahasiswa/:id", controller.DeleteMahasiswa)

	return r
}
