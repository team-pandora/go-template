package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/team-pandora/go-template/server/errors"
	"github.com/team-pandora/go-template/server/feature"
	"github.com/team-pandora/go-template/utils"
)

// useServerRouter registers the server router with the provided gin router.
func useServerRouter(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		utils.GinAbortWithError(c, errors.InvalidRouteError)
	})

	router.Any("/isAlive", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	apiRouter := router.Group("/api")

	featureRouter := apiRouter.Group("/feature")
	feature.RegisterRotuer(featureRouter)
}
