package server

import (
	"github.com/MichaelSimkin/go-template/server/errors"
	"github.com/MichaelSimkin/go-template/server/feature"
	"github.com/MichaelSimkin/go-template/utils"
	"github.com/gin-gonic/gin"
)

// useServerRouter registers the server router with the provided gin router.
func useServerRouter(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		utils.GinAbortWithError(c, errors.InvalidRouteError)
	})

	router.Any("/isAlive", func(c *gin.Context) {
		c.String(200, "OK")
	})

	apiRouter := router.Group("/api")

	featureRouter := apiRouter.Group("/feature")
	feature.RegisterRotuer(featureRouter)
}
