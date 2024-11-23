package controllers

import (
	"fmt"
	"net/http"

	"github.com/attanabilrabbani/go-crud/initializers"
	"github.com/attanabilrabbani/go-crud/models"
	"github.com/gin-gonic/gin"
)

func MahasiswaAdd(c *gin.Context) {
	//get req from body
	var reqbody struct {
		Nama     string
		NPM      string
		Semester int
		Ipk      float64
	}

	c.Bind(&reqbody)

	//create
	mahasiswa := models.Mahasiswa{
		Nama:     reqbody.Nama,
		NPM:      reqbody.NPM,
		Semester: reqbody.Semester,
		Ipk:      reqbody.Ipk,
	}

	err := initializers.DB.Create(&mahasiswa).Error

	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, mahasiswa)
}

func MahasiswaGet(c *gin.Context) {
	var mahasiswa []models.Mahasiswa
	initializers.DB.Find(&mahasiswa)

	c.JSON(http.StatusOK, mahasiswa)

}

func MahasiswaGetById(c *gin.Context) {
	//get id from url
	mahasiswaId := c.Param("id")

	var mahasiswa models.Mahasiswa
	initializers.DB.First(&mahasiswa, mahasiswaId)

	c.JSON(http.StatusOK, mahasiswa)
}

func MahasiswaUpdate(c *gin.Context) {
	mahasiswaId := c.Param("id")

	var updatebody struct {
		Nama     string
		NPM      string
		Semester *int
		Ipk      *float64
	}

	c.Bind(&updatebody)

	var mahasiswa models.Mahasiswa
	initializers.DB.First(&mahasiswa, mahasiswaId)

	//update
	updateData := make(map[string]interface{})

	if updatebody.Nama != "" {
		updateData["Nama"] = updatebody.Nama
	}
	if updatebody.NPM != "" {
		updateData["NPM"] = updatebody.NPM
	}
	if updatebody.Semester != nil {
		updateData["Semester"] = updatebody.Semester
	}
	if updatebody.Ipk != nil {
		updateData["Ipk"] = updatebody.Ipk
	}

	initializers.DB.Model(&mahasiswa).Updates(updateData)

	c.JSON(http.StatusOK, updateData)
}

func MahasiswaDelete(c *gin.Context) {
	mahasiswaId := c.Param("id")

	initializers.DB.Delete(&models.Mahasiswa{}, mahasiswaId)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Mahasiswa dengan ID %s berhasil dihapus", mahasiswaId),
	})
}
