package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Nombre        string `json:"nombre"`
	Paterno       string `json:"paterno"`
	Materno       string `json:"materno"`
	Correo        string `json:"correo"`
	Edad          string `json:"edad"`
	Telefono      string `json:"telefono"`
	Direccion     string `json:"direccion"`
	DNI           string `json:"dni"`
	Estado        string `json:"estado"`
	Inscripcion   []Inscripcion
	RolPersona    []RolPersona
	Trabajo       []Trabajo
	PlanAcademico []PlanAcademico
}
