package rutas

import (
	"api_apk_votaciones/models"

	"github.com/gin-gonic/gin"
)

func RegisterPostulanteRoutes(r *gin.Engine) {
	r.GET("/postulantes", func(c *gin.Context) {
		postulante := models.Postulante{}
		postulantes, err := postulante.GetAllPostulantes()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, postulantes)
	})

	r.POST("/postulantes", func(c *gin.Context) {
		var postulante models.Postulante
		if err := c.ShouldBindJSON(&postulante); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err := postulante.CreatePostulante(&postulante)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, postulante)
	})

	r.PUT("/postulantes/:id", func(c *gin.Context) {
		var postulante *models.Postulante
		id := c.Param("id")

		postulante, err := postulante.GetPostulanteByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "Postulante no encontrado"})
			return
		}

		if err := c.ShouldBindJSON(&postulante); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err = postulante.UpdatePostulante(postulante)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, postulante)
	})
}
