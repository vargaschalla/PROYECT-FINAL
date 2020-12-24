package models

import "gorm.io/gorm"

type Trabajo struct {
	gorm.Model
	Nombre       string `json:"nombre"`
	FechaInicio  string `json:"fechainicio"`
	FechaFin     string `json:"fechafin"`
	FechaEntrega string `json:"fechaentrega"`
	Nota         string `json:"nota"`
	Estado       string `json:"estado"`
	PersonID     string `gorm:"size:191"`
	RecursoID    string `gorm:"size:191"`
	SecuenciaID  string `gorm:"size:191"`
	Person       Person
	Recurso      Recurso
	Secuencia    Secuencia
}
