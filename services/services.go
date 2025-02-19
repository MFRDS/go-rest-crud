package services

import (
	"go-rest-crud/models"
	"go-rest-crud/repository"
)

type MahasiswaService interface {
	GetAll() ([]models.Mahasiswa, error)
	GetByID(id int) (models.Mahasiswa, error)
	Create(mahasiswa models.Mahasiswa) error
	Update(mahasiswa models.Mahasiswa) error
	Delete(id int) error
}

type mahasiswaService struct {
	repo repository.MahasiswaRepository
}

func NewMahasiswaService(repo repository.MahasiswaRepository) MahasiswaService {
	return &mahasiswaService{repo: repo}
}

func (s *mahasiswaService) GetAll() ([]models.Mahasiswa, error) {
	return s.repo.GetAll()
}

func (s *mahasiswaService) GetByID(id int) (models.Mahasiswa, error) {
	return s.repo.GetByID(id)
}

func (s *mahasiswaService) Create(mahasiswa models.Mahasiswa) error {
	return s.repo.Create(mahasiswa)
}

func (s *mahasiswaService) Update(mahasiswa models.Mahasiswa) error {
	return s.repo.Update(mahasiswa)
}

func (s *mahasiswaService) Delete(id int) error {
	return s.repo.Delete(id)
}
