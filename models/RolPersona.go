package models

import "gorm.io/gorm"

type RolPersona struct {
	gorm.Model
	PersonID string `gorm:"size:191"`
	RolID    string `gorm:"size:191"`
	Person   Person
	Rol      Rol
}
