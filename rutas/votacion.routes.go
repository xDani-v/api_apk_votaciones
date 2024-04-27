// rutas/usuario.routes.go
package rutas

import (
	"api_apk_votaciones/models"

	"github.com/gin-gonic/gin"
)

func RegisterVotacionRoutes(r *gin.Engine) {
	// ... tus rutas de usuario existentes ...

	r.GET("/votaciones", func(c *gin.Context) {
		votacion := models.Votacion{}
		votaciones, err := votacion.GetAllVotacions()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, votaciones)
	})

	r.POST("/votaciones", func(c *gin.Context) {
		var votacion models.Votacion
		if err := c.ShouldBindJSON(&votacion); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err := votacion.CreateVotacion(&votacion)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, votacion)
	})

	r.PUT("/votaciones/:id", func(c *gin.Context) {
		var votacion *models.Votacion
		id := c.Param("id")

		votacion, err := votacion.GetVotacionByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "Votacion no encontrada"})
			return
		}

		if err := c.ShouldBindJSON(&votacion); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err = votacion.UpdateVotacion(votacion)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, votacion)
	})
}
