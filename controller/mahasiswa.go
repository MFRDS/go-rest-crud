package controller

import (
	"fmt"
	"net/http"

	"go-rest-crud/database"
	"go-rest-crud/models"

	"github.com/gin-gonic/gin"
)

func GetMahasiswa(c *gin.Context) {
	var mahasiswa []models.Mahasiswa
	database.DB.Find(&mahasiswa)
	c.JSON(http.StatusOK, mahasiswa)
}

func CreateMahasiswa(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	c.BindJSON(&mahasiswa)
	err := database.DB.Create(&mahasiswa)
	if err.Error != nil {
		fmt.Println(err.Error)
	}
	c.JSON(http.StatusCreated, mahasiswa)
}

func GetMahasiswaByID(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	id := c.Param("id")
	if err := database.DB.Where("id = ?", id).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data mahasiswa tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, mahasiswa)
}

func UpdateMahasiswa(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	id := c.Param("id")
	if err := database.DB.Where("id = ?", id).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data mahasiswa tidak ditemukan"})
		return
	}
	c.BindJSON(&mahasiswa)
	database.DB.Save(&mahasiswa)
	c.JSON(http.StatusOK, mahasiswa)
}

func DeleteMahasiswa(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	id := c.Param("id")
	if err := database.DB.Where("id = ?", id).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data mahasiswa tidak ditemukan"})
		return
	}
	database.DB.Delete(&mahasiswa)
	c.JSON(http.StatusOK, gin.H{"message": "Data mahasiswa berhasil dihapus"})
}
