package controllers

import (
	"net/http"
	"strconv"

	"go-rest-crud/models"
	"go-rest-crud/services"

	"github.com/gin-gonic/gin"
)

type MahasiswaController struct {
	service services.MahasiswaService
}

func NewMahasiswaController(service services.MahasiswaService) *MahasiswaController {
	return &MahasiswaController{service: service}
}

func (c *MahasiswaController) GetAll(ctx *gin.Context) {
	mahasiswa, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, mahasiswa)
}

func (c *MahasiswaController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	mahasiswa, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa not found"})
		return
	}

	ctx.JSON(http.StatusOK, mahasiswa)
}

func (c *MahasiswaController) Create(ctx *gin.Context) {
	var mahasiswa models.Mahasiswa
	if err := ctx.ShouldBindJSON(&mahasiswa); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Create(mahasiswa); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Mahasiswa created successfully"})
}

func (c *MahasiswaController) Update(ctx *gin.Context) {
	var mahasiswa models.Mahasiswa
	if err := ctx.ShouldBindJSON(&mahasiswa); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Update(mahasiswa); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Mahasiswa updated successfully"})
}

func (c *MahasiswaController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.service.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Mahasiswa deleted successfully"})
}
