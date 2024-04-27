package models

import (
	"api_apk_votaciones/conexion"

	"time"
)

type Votacion struct {
	IDUsuario   uint
	Nombre      string
	FechaInicio time.Time
	FechaFin    time.Time
	Codigo      string
	Estado      string
}

func (v *Votacion) GetAllVotacions() ([]Votacion, error) {
	var votacions []Votacion
	err := conexion.Db.Find(&votacions).Error
	return votacions, err
}

func (v *Votacion) CreateVotacion(votacion *Votacion) error {
	err := conexion.Db.Create(votacion).Error
	return err
}

func (v *Votacion) UpdateVotacion(votacion *Votacion) error {
	err := conexion.Db.Save(votacion).Error
	return err
}

func (v *Votacion) GetVotacionByID(id string) (*Votacion, error) {
	var votacion Votacion
	err := conexion.Db.First(&votacion, id).Error
	if err != nil {
		return nil, err
	}
	return &votacion, nil
}
