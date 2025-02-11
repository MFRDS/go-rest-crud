package models

type Mahasiswa struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Nama  string `json:"nama"`
	Nim   string `json:"nim"`
	Prodi string `json:"prodi"`
}

func (Mahasiswa) TableName() string {
	return "mahasiswa"
}
