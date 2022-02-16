package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// GinAbortWithError is a helper function to abort the gin context with the provided error.
func GinAbortWithError(c *gin.Context, err error) {
	if err := c.Error(err); err != nil {
		fmt.Println("Gin error:", err)
	}
	c.Abort()
}
