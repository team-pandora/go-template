package errors

import (
	"net/http"
)

// FeatureError is an error that is used to test error handling.
var FeatureError = ServerError{Code: http.StatusInternalServerError, Message: "Something went wrong"}

// NewInvalidRequestError is used to return an error when the server fails to decode a request.
// The error is returned in the meta.
func NewInvalidRequestError(err error) ServerError {
	return ServerError{Code: http.StatusBadRequest, Message: "Invalid request", Meta: ErrorMeta{"error": err.Error()}}
}

// DuplicateKeyError is used when a mongo duplicate key error occurs.
var DuplicateKeyError = ServerError{Code: http.StatusConflict, Message: "Document already exists"}

// InvalidMongoIDError is used when a returned mongo ID is invalid.
var InvalidMongoIDError = ServerError{Code: http.StatusConflict, Message: "Mongo returned invalid id"}

// NewUnknownMongoError creates a new unknown mongo error with the provided error in the meta.
func NewUnknownMongoError(err error) ServerError {
	return ServerError{Code: http.StatusInternalServerError, Message: "Unknown mongo error", Meta: ErrorMeta{"error": err.Error()}}
}

// NewInvalidFiltersError is used to return an error when the filters provided to the mongo find method are invalid.
// The error is returned in the meta.
func NewInvalidFiltersError(err error) ServerError {
	return ServerError{Code: http.StatusBadRequest, Message: "Invalid Filters", Meta: ErrorMeta{"error": err.Error()}}
}

// NewFailedToDecodeError is used to return an error when the server fails to decode a mongo document.
// The error is returned in the meta.
func NewFailedToDecodeError(err error) ServerError {
	return ServerError{Code: http.StatusInternalServerError, Message: "Failed to decode mongo document", Meta: ErrorMeta{"error": err.Error()}}
}

// NewMongoCursorError is used to return an error when the server fails to iterate over a mongo cursor.
// The mongo error is returned in the meta.
func NewMongoCursorError(err error) ServerError {
	return ServerError{Code: http.StatusInternalServerError, Message: "Mongo cursor error", Meta: ErrorMeta{"error": err.Error()}}
}
