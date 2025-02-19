package routes

import (
	"go-rest-crud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mahasiswaController *controllers.MahasiswaController) *gin.Engine {
	r := gin.Default()

	mahasiswaRoutes := r.Group("/mahasiswa")
	{
		mahasiswaRoutes.GET("", mahasiswaController.GetAll)
		mahasiswaRoutes.GET("/:id", mahasiswaController.GetByID)
		mahasiswaRoutes.POST("", mahasiswaController.Create)
		mahasiswaRoutes.PUT("", mahasiswaController.Update)
		mahasiswaRoutes.DELETE("/:id", mahasiswaController.Delete)
	}

	return r
}
