package feature

import (
	"github.com/gin-gonic/gin"
	"github.com/team-pandora/go-template/server/errors"
	"github.com/team-pandora/go-template/utils"
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
