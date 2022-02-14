package feature

import (
	"net/http"

	"github.com/MichaelSimkin/go-template/server/errors"
	"github.com/gin-gonic/gin"
)

var Service IService = &service{}

type service struct{}

type IService interface {
	createDocumet(c *gin.Context)
	getDocumets(c *gin.Context)
}

func (service) createDocumet(c *gin.Context) {
	// Get the request document
	var document = &FeatureModel{}
	err := c.ShouldBindJSON(document)
	if err != nil {
		c.Error(errors.NewInvalidRequestError(err))
		c.Abort()
		return
	}

	// Validate the request document
	err = Validate.Struct(document)
	if err != nil {
		c.Error(errors.NewInvalidRequestError(err))
		c.Abort()
		return
	}

	// Set document timestamps
	document.setTimestamps()

	// Create the document
	result, err := Repository.createDocument(c.Request.Context(), *document)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (service) getDocumets(c *gin.Context) {
	filters, ok := c.GetQuery("filters")
	if !ok {
		filters = "{}"
	}
	result, err := Repository.getDocuments(c.Request.Context(), []byte(filters))
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result)
}
