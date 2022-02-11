package errors

import (
	"net/http"
)

var FeatureError = &ServerError{Code: http.StatusInternalServerError, Message: "Something went wrong", Meta: ErrorMeta{"origin": "feature"}}
