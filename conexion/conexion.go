// conexion.go
package conexion

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=sistema_votacion sslmode=disable password=root")
	if err != nil {
		panic("failed to connect database")
	}
}
