package main

import (
	"github.com/attanabilrabbani/go-crud/initializers"
	"github.com/attanabilrabbani/go-crud/models"
	"github.com/attanabilrabbani/go-crud/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Mahasiswa{})
	r := gin.Default()

	routes.MahasiswaRoutes(r)

	r.Run()

}
