package main

import (
	"api_apk_votaciones/rutas"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	rutas.RegisterUsuarioRoutes(r)
	rutas.RegisterVotacionRoutes(r)
	rutas.RegisterPostulanteRoutes(r)
	rutas.RegisterVotoRoutes(r)
	r.Run()
}
