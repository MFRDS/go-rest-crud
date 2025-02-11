package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/MFRDS/go-rest-crud/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root@tcp(localhost:3306)/db_mahasiswa"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	database.AutoMigrate(&models.Mahasiswa{})
	DB = database

}
