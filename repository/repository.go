package repository

import (
	"go-rest-crud/database"
	"go-rest-crud/models"
)

func GetAllMahasiswa() ([]models.Mahasiswa, error) {
	var mahasiswa []models.Mahasiswa
	result := database.DB.Find(&mahasiswa)
	return mahasiswa, result.Error
}

func GetMahasiswaByID(id string) (models.Mahasiswa, error) {
	var mahasiswa models.Mahasiswa
	result := database.DB.Where("id = ?", id).First(&mahasiswa)
	return mahasiswa, result.Error
}

func CreateMahasiswa(mahasiswa *models.Mahasiswa) error {
	return database.DB.Create(mahasiswa).Error
}

func UpdateMahasiswa(mahasiswa *models.Mahasiswa) error {
	return database.DB.Save(mahasiswa).Error
}

func DeleteMahasiswa(id string) error {
	var mahasiswa models.Mahasiswa
	if err := database.DB.Where("id = ?", id).First(&mahasiswa).Error; err != nil {
		return err
	}
	return database.DB.Delete(&mahasiswa).Error
}
