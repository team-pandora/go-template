package feature

import (
	"github.com/MichaelSimkin/go-template/server/errors"
	"github.com/MichaelSimkin/go-template/utils"
	"github.com/gin-gonic/gin"
)

// RegisterRotuer registers the feature router with the provided gin router.
func RegisterRotuer(router *gin.RouterGroup) {
	router.POST("/", Service.CreateDocumet)
	router.GET("/", Service.GetDocumets)

	// Tests for error handling
	router.GET("/error", func(c *gin.Context) {
		utils.GinAbortWithError(c, errors.FeatureError)
	})
}
