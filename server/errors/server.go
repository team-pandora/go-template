package errors

import (
	"fmt"
)

type ErrorMeta = map[string]interface{}

type ServerError struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Meta    ErrorMeta `json:"meta,omitempty"`
}

func (e ServerError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
