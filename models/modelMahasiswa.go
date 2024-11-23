package models

import "gorm.io/gorm"

type Mahasiswa struct {
	gorm.Model
	Nama     string
	NPM      string
	Semester int
	Ipk      float64
}
