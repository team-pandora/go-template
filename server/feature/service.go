package feature

import (
	"net/http"

	"github.com/MichaelSimkin/go-template/server/errors"
	"github.com/gin-gonic/gin"
)

var service *featureService

type featureService struct{}

// initService initializes the service, validate and repository global variables.
func initService() {
	initValidator()
	initRepository()
	service = &featureService{}
}

func (featureService) createDocumet(c *gin.Context) {
	// get the request document
	var document = &featureModel{}
	err := c.ShouldBindJSON(document)
	if err != nil {
		c.Error(errors.NewInvalidRequestError(err))
		c.Abort()
		return
	}

	// validate the request document
	err = validate.Struct(document)
	if err != nil {
		c.Error(errors.NewInvalidRequestError(err))
		c.Abort()
		return
	}

	// set document timestamps
	document.setTimestamps()

	// create the document
	result, err := repository.createDocument(c.Request.Context(), *document)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (featureService) getDocumets(c *gin.Context) {
	filters, ok := c.GetQuery("filters")
	if !ok {
		filters = "{}"
	}
	result, err := repository.getDocuments(c.Request.Context(), []byte(filters))
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result)
}
