package server

import (
	"github.com/MichaelSimkin/go-template/server/feature"
	"github.com/gin-gonic/gin"
)

func UseServerRouter(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.String(404, "Invalid Route")
	})

	router.Any("/isAlive", func(c *gin.Context) {
		c.String(200, "OK")
	})

	apiRouter := router.Group("/api")

	featureRouter := apiRouter.Group("/feature")
	feature.UseFeatureRotuer(featureRouter)
}
