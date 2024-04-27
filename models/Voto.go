package models

import (
	"api_apk_votaciones/conexion"

	"github.com/jinzhu/gorm"
)

type Voto struct {
	gorm.Model
	IDUsuario    uint
	IDVotacion   uint
	IDPostulante uint
	Voto         string `gorm:"type:varchar(15)"`
	Estado       string `gorm:"type:varchar(20)"`
}

func (v *Voto) GetAllVotos() ([]Voto, error) {
	var votos []Voto
	err := conexion.Db.Find(&votos).Error
	return votos, err
}

func (v *Voto) CreateVoto(voto *Voto) error {
	err := conexion.Db.Create(voto).Error
	return err
}

func (v *Voto) UpdateVoto(voto *Voto) error {
	err := conexion.Db.Save(voto).Error
	return err
}

func (v *Voto) GetVotoByID(id string) (*Voto, error) {
	var voto Voto
	err := conexion.Db.First(&voto, id).Error
	if err != nil {
		return nil, err
	}
	return &voto, nil
}
