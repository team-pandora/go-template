package feature

import (
	"net/http"

	"github.com/MichaelSimkin/go-template/server/errors"
	"github.com/gin-gonic/gin"
)

// UseFeatureRouter registers the feature router with the provided gin router.
func UseFeatureRotuer(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Feature Works!")
	})

	router.GET("/error", func(c *gin.Context) {
		c.Error(errors.FeatureError)
		c.Abort()
	})

	router.GET("/multipleErrors", func(c *gin.Context) {
		c.Error(errors.FeatureError)
		c.Error(errors.FeatureError)
		c.Error(errors.FeatureError)
		c.Abort()
	})
}
