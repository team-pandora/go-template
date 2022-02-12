package server

import (
	"fmt"
	"net/http"

	"github.com/MichaelSimkin/go-template/server/errors"
	"github.com/gin-gonic/gin"
)

// errorHandler executes all the middlewares and then checks for errors.
// If an error is found, it is parsed into a ServerError and sent to the client.
// If multiple errors are found, they are all parsed into ServerErrors and sent to the client as a single ServerError that contains all errors in the Meta field.
func errorHandler(c *gin.Context) {
	// execute all middleware logic
	c.Next()

	if len(c.Errors) == 0 {
		return
	}

	var parsedErrors = make([]*errors.ServerError, 0, len(c.Errors))
	for _, err := range c.Errors {
		parsedErrors = append(parsedErrors, parseError(err.Err))
	}

	if len(parsedErrors) == 1 {
		c.JSON(parsedErrors[0].Code, parsedErrors[0])
	} else {
		c.JSON(http.StatusInternalServerError, &errors.ServerError{
			Code:    http.StatusInternalServerError,
			Message: "Multiple errors occurred",
			Meta:    errors.ErrorMeta{"errors": parsedErrors},
		})
	}
}

// parseError parses an error into a ServerError.
// If the error is not a predefined error, parse it into a ServerError.
func parseError(err error) *errors.ServerError {
	switch err := err.(type) {
	case errors.ServerError:
		return &err
	default:
		return &errors.ServerError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Internal Server Error: %v", err),
		}
	}
}
