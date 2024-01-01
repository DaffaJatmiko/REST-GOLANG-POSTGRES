package models

import (
	"gorm.io/gorm"
	"time"
)

type Mahasiswa struct {
	ID              uint   `gorm:"primary key; increment" json:"id"`
	Name            string `json:"name"`
	BirthDate       time.Time `json:"birth_date"`
	Gender          string `json:"gender"`
	IDJurusan       uint `json:"id_jurusan"`
	IDHobi          uint `json:"id_hobi"`
	MahasiswaJurusan string `json:"mahasiswa_jurusan"`
	Jurusan         Jurusan `gorm:"foreignKey:IDJurusan"`
	Hobi            Hobi    `gorm:"foreignKey:IDHobi"`
}

type Jurusan struct {
	ID    uint   `gorm:"primary key; increment" json:"id"`
	Nama  string `json:"nama"`
}

type Hobi struct {
	ID    uint   `gorm:"primary key; increment" json:"id"`
	Nama  string `json:"nama"`
}

func MigrateMahasiswa(db *gorm.DB) error {
	// AutoMigrate untuk Mahasiswa, Jurusan, dan Hobi
	err := db.AutoMigrate(&Mahasiswa{}, &Jurusan{}, &Hobi{})
	if err != nil {
		return err
	}

	// Mengeksekusi pernyataan SQL untuk menambahkan kunci asing ke Jurusan
	err = db.Exec("ALTER TABLE mahasiswas ADD CONSTRAINT fk_jurusan FOREIGN KEY (id_jurusan) REFERENCES jurusans(id)").Error
	if err != nil {
		return err
	}

	// Mengeksekusi pernyataan SQL untuk menambahkan kunci asing ke Hobi
	err = db.Exec("ALTER TABLE mahasiswas ADD CONSTRAINT fk_hobi FOREIGN KEY (id_hobi) REFERENCES hobis(id)").Error
	if err != nil {
		return err
	}

	return nil
}


