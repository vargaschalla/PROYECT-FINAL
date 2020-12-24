package models

import "gorm.io/gorm"

type Sesion struct {
	gorm.Model
	Titulo   string `json:"titulo"`
	Tema     string `json:"tema"`
	Estado   string `json:"estado"`
	UnidadID string `gorm:"size:191"`
	Recurso  []Recurso
	Unidad   Unidad
}
