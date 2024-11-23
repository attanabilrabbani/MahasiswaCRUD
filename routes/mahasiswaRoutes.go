package routes

import (
	"github.com/attanabilrabbani/go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func MahasiswaRoutes(r *gin.Engine) {

	r.POST("/mahasiswa/add/", controllers.MahasiswaAdd)
	r.GET("/mahasiswa/", controllers.MahasiswaGet)
	r.GET("/mahasiswa/:id", controllers.MahasiswaGetById)
	r.PUT("/mahasiswa/:id", controllers.MahasiswaUpdate)
	r.DELETE("/mahasiswa/:id", controllers.MahasiswaDelete)
}
