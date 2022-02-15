package errors

import (
	"fmt"
)

// ErrorMeta is a map of key value pairs that can be used to add additional information to an error.
type ErrorMeta = map[string]interface{}

// ServerError is a generic error that is used to return an error to the client.
type ServerError struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Meta    ErrorMeta `json:"meta,omitempty"`
}

func (e ServerError) Error() string {
	if len(e.Meta) == 0 {
		return fmt.Sprintf("%v.", e.Message)
	}
	return fmt.Sprintf("%v, %v.", e.Message, e.Meta)
}

// InvalidRouteError is used to return an error when the route is not found.
var InvalidRouteError = ServerError{Code: 400, Message: "Invalid route"}
