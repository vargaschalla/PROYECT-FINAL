package models

import "gorm.io/gorm"

type Recurso struct {
	gorm.Model
	Titulo        string `json:"titulo"`
	Descripcion   string `json:"descripcion"`
	Estado        string `json:"estado"`
	SesionID      string `gorm:"size:191"`
	TipoRecursoID string `gorm:"size:191"`
	Trabajo       []Trabajo
	Sesion        Sesion
	TipoRecurso   TipoRecurso
}
