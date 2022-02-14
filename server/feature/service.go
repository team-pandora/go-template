package feature

import (
	"net/http"

	"github.com/MichaelSimkin/go-template/server/errors"
	"github.com/gin-gonic/gin"
)

var Service = &service{}

type service struct{}

func (service) createDocumet(c *gin.Context) {
	// Get the request body and validate it
	var document = &FeatureModel{}
	if !getRequestBody(c, document) {
		return
	}

	// Set document timestamps
	document.setTimestamps()

	// Create the document in the database
	result, err := Repository.createDocument(c.Request.Context(), *document)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	// Return the created document
	c.JSON(http.StatusCreated, result)
}

func (service) getDocumets(c *gin.Context) {
	// Get filters from query string or set to default empty filters
	filters, ok := c.GetQuery("filters")
	if !ok {
		filters = "{}"
	}

	// Get the documents from the database
	result, err := Repository.getDocuments(c.Request.Context(), filters)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result)
}

// getRequestBody binds and validates the request body from the provided gin context.
func getRequestBody(c *gin.Context, body interface{}) bool {
	if bindRequestBody(c, body) && validateRequestBody(c, body) {
		return true
	}
	return false
}

// bindRequestBody binds the request body from the provided gin context to the provided body.
func bindRequestBody(c *gin.Context, body interface{}) bool {
	err := c.ShouldBindJSON(body)
	if err != nil {
		c.Error(errors.NewInvalidRequestError(err))
		c.Abort()
		return false
	}
	return true
}

// validateRequestBody validates the request body from the provided gin context.
func validateRequestBody(c *gin.Context, body interface{}) bool {
	err := Validate.Struct(body)
	if err != nil {
		c.Error(errors.NewInvalidRequestError(err))
		c.Abort()
		return false
	}
	return true
}
