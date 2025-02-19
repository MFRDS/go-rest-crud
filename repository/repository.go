package repository

import (
	"go-rest-crud/database"
	"go-rest-crud/models"
)

type MahasiswaRepository interface {
	GetAll() ([]models.Mahasiswa, error)
	GetByID(id int) (models.Mahasiswa, error)
	Create(mahasiswa models.Mahasiswa) error
	Update(mahasiswa models.Mahasiswa) error
	Delete(id int) error
}

type mahasiswaRepository struct{}

func NewMahasiswaRepository() MahasiswaRepository {
	return &mahasiswaRepository{}
}

func (r *mahasiswaRepository) GetAll() ([]models.Mahasiswa, error) {
	var mahasiswa []models.Mahasiswa
	err := database.DB.Find(&mahasiswa).Error
	return mahasiswa, err
}

func (r *mahasiswaRepository) GetByID(id int) (models.Mahasiswa, error) {
	var mahasiswa models.Mahasiswa
	err := database.DB.First(&mahasiswa, id).Error
	return mahasiswa, err
}

func (r *mahasiswaRepository) Create(mahasiswa models.Mahasiswa) error {
	return database.DB.Create(&mahasiswa).Error
}

func (r *mahasiswaRepository) Update(mahasiswa models.Mahasiswa) error {
	return database.DB.Save(&mahasiswa).Error
}

func (r *mahasiswaRepository) Delete(id int) error {
	return database.DB.Delete(&models.Mahasiswa{}, id).Error
}
