package services

import (
	"go-rest-crud/models"
	"go-rest-crud/repository"
)

func GetMahasiswaList() ([]models.Mahasiswa, error) {
	return repository.GetAllMahasiswa()
}

func GetMahasiswaDetail(id string) (models.Mahasiswa, error) {
	return repository.GetMahasiswaByID(id)
}

func AddMahasiswa(mahasiswa *models.Mahasiswa) error {
	return repository.CreateMahasiswa(mahasiswa)
}

func ModifyMahasiswa(mahasiswa *models.Mahasiswa) error {
	return repository.UpdateMahasiswa(mahasiswa)
}

func RemoveMahasiswa(id string) error {
	return repository.DeleteMahasiswa(id)
}
