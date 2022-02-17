package utils

import (
	"github.com/gin-gonic/gin"
)

// GinAbortWithError is a helper function to abort the gin context with the provided error.
func GinAbortWithError(c *gin.Context, err error) {
	if ginErr := c.Error(err); ginErr != nil {
		Log.Errorf("error while aborting gin context: %v", ginErr)
	}
	c.Abort()
}
