package rutas

import (
	"api_apk_votaciones/models"

	"github.com/gin-gonic/gin"
)

func RegisterVotoRoutes(r *gin.Engine) {
	r.GET("/votos", func(c *gin.Context) {
		voto := models.Voto{}
		votos, err := voto.GetAllVotos()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, votos)
	})

	r.POST("/votos", func(c *gin.Context) {
		var voto models.Voto
		if err := c.ShouldBindJSON(&voto); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err := voto.CreateVoto(&voto)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, voto)
	})

	r.PUT("/votos/:id", func(c *gin.Context) {
		var voto *models.Voto
		id := c.Param("id")

		voto, err := voto.GetVotoByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "Voto no encontrado"})
			return
		}

		if err := c.ShouldBindJSON(&voto); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err = voto.UpdateVoto(voto)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, voto)
	})
}
