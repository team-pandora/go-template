package utils

import (
	"github.com/gin-gonic/gin"
)

// GinAbortWithError is a helper function to abort the gin context with the provided error.
func GinAbortWithError(c *gin.Context, err error) {
	if err := c.Error(err); err != nil {
		Log.Errorf("Error while aborting gin context: %v", err)
	}
	c.Abort()
}
