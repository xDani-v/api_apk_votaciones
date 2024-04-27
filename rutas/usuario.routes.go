// rutas/usuario.routes.go
package rutas

import (
	"api_apk_votaciones/models"

	"github.com/gin-gonic/gin"
)

func RegisterUsuarioRoutes(r *gin.Engine) {
	r.GET("/usuarios", func(c *gin.Context) {
		usuario := models.Usuario{}
		usuarios, err := usuario.GetAllUsuarios()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, usuarios)
	})

	r.POST("/usuarios", func(c *gin.Context) {
		var usuario models.Usuario
		if err := c.ShouldBindJSON(&usuario); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err := usuario.CreateUsuario(&usuario)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, usuario)
	})

	r.PUT("/usuarios/:id", func(c *gin.Context) {
		var usuario *models.Usuario
		id := c.Param("id")

		usuario, err := usuario.GetUsuarioByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "Usuario no encontrado"})
			return
		}

		if err := c.ShouldBindJSON(usuario); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err = usuario.UpdateUsuario(usuario)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, usuario)
	})

	r.POST("/usuarios/login", func(c *gin.Context) {
		var loginInfo struct {
			Identificacion string `json:"identificacion"`
			Password       string `json:"password"`
		}
		if err := c.ShouldBindJSON(&loginInfo); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		usuario := models.Usuario{}
		user, err := usuario.GetUsuarioByIdentificacionAndPassword(loginInfo.Identificacion, loginInfo.Password)
		if err != nil {
			c.JSON(404, gin.H{"error": "Usuario no encontrado"})
			return
		}
		c.JSON(200, user)
	})
}
