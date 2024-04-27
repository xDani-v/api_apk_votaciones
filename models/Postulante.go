package models

import (
	"api_apk_votaciones/conexion"

	"github.com/jinzhu/gorm"
)

type Postulante struct {
	gorm.Model
	IDVotacion  uint
	Nombre      string `gorm:"type:varchar(50)"`
	Grupo       string `gorm:"type:varchar(50)"`
	Descripcion string `gorm:"type:varchar(100)"`
	Estado      string `gorm:"type:varchar(20)"`
}

func (p *Postulante) GetAllPostulantes() ([]Postulante, error) {
	var postulantes []Postulante
	err := conexion.Db.Find(&postulantes).Error
	return postulantes, err
}

func (p *Postulante) CreatePostulante(postulante *Postulante) error {
	err := conexion.Db.Create(postulante).Error
	return err
}

func (p *Postulante) UpdatePostulante(postulante *Postulante) error {
	err := conexion.Db.Save(postulante).Error
	return err
}

func (p *Postulante) GetPostulanteByID(id string) (*Postulante, error) {
	var postulante Postulante
	err := conexion.Db.First(&postulante, id).Error
	if err != nil {
		return nil, err
	}
	return &postulante, nil
}
