package models

import "gorm.io/gorm"

type PlanAcademico struct {
	gorm.Model
	Nombre      string `json:"nombre"`
	Estado      string `json:"estado"`
	Inscripcion []Inscripcion
	Unidad      []Unidad
	CursoID     string `gorm:"size:191"`
	PeriodoID   string `gorm:"size:191"`
	PersonID    string `gorm:"size:191"`
	SeccionID   string `gorm:"size:191"`
	Curso       Curso
	Periodo     Periodo
	Person      Person
	Seccion     Seccion
}
