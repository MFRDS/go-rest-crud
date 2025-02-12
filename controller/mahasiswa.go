package controller

import (
	"go-rest-crud/models"
	"go-rest-crud/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMahasiswa(c *gin.Context) {
	mahasiswa, err := services.GetMahasiswaList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan data mahasiswa"})
		return
	}
	c.JSON(http.StatusOK, mahasiswa)
}

func GetMahasiswaByID(c *gin.Context) {
	id := c.Param("id")
	mahasiswa, err := services.GetMahasiswaDetail(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, mahasiswa)
}

func CreateMahasiswa(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	if err := services.AddMahasiswa(&mahasiswa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan mahasiswa"})
		return
	}
	c.JSON(http.StatusCreated, mahasiswa)
}

func UpdateMahasiswa(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	if err := services.ModifyMahasiswa(&mahasiswa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui mahasiswa"})
		return
	}
	c.JSON(http.StatusOK, mahasiswa)
}

func DeleteMahasiswa(c *gin.Context) {
	id := c.Param("id")
	if err := services.RemoveMahasiswa(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus mahasiswa"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mahasiswa berhasil dihapus"})
}
