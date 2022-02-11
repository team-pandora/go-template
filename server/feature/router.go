package feature

import "github.com/gin-gonic/gin"

func UseFeatureRotuer(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Feature Works!")
	})
}
