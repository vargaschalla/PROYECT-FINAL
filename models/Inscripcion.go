package models

import "gorm.io/gorm"

type Inscripcion struct {
	gorm.Model
	PlanAcademicoID string `gorm:"size:191"`
	PersonID        string `gorm:"size:191"`
	PlanAcademico   PlanAcademico
	Person          Person
}
