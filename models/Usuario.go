// usuario.go
package models

import (
	"api_apk_votaciones/conexion"

	"github.com/jinzhu/gorm"
)

type Usuario struct {
	gorm.Model
	Identificacion string `gorm:"type:varchar(20)"`
	NombreCompleto string `gorm:"type:varchar(100)"`
	Correo         string `gorm:"type:varchar(70);unique_index"`
	Rol            string `gorm:"type:varchar(20)"`
	Password       string `gorm:"type:varchar(30)"`
	Estado         string `gorm:"type:varchar(15)"`
}

func (u *Usuario) GetAllUsuarios() ([]Usuario, error) {
	var usuarios []Usuario
	err := conexion.Db.Find(&usuarios).Error
	return usuarios, err
}

func (u *Usuario) CreateUsuario(usuario *Usuario) error {
	err := conexion.Db.Create(usuario).Error
	return err
}

func (u *Usuario) UpdateUsuario(usuario *Usuario) error {
	err := conexion.Db.Save(usuario).Error
	return err
}

func (u *Usuario) GetUsuarioByID(id string) (*Usuario, error) {
	var usuario Usuario
	err := conexion.Db.First(&usuario, id).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (u *Usuario) GetUsuarioByIdentificacionAndPassword(identificacion string, password string) (*Usuario, error) {
	var usuario Usuario
	err := conexion.Db.Where("identificacion = ? AND password = ?", identificacion, password).First(&usuario).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}
