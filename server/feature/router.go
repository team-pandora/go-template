package feature

import (
	"github.com/MichaelSimkin/go-template/server/errors"
	"github.com/gin-gonic/gin"
)

// RegisterRotuer registers the feature router with the provided gin router.
func RegisterRotuer(router *gin.RouterGroup) {
	router.POST("/", Service.createDocumet)
	router.GET("/", Service.getDocumets)

	// tests for error handling
	router.GET("/error", func(c *gin.Context) {
		c.Error(errors.FeatureError)
		c.Abort()
	})
	// tests for error handling
	router.GET("/multipleErrors", func(c *gin.Context) {
		c.Error(errors.FeatureError)
		c.Error(errors.FeatureError)
		c.Error(errors.FeatureError)
		c.Abort()
	})
}
