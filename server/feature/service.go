package feature

import (
	"net/http"

	"github.com/MichaelSimkin/go-template/server/errors"
	"github.com/gin-gonic/gin"
)

var service = &featureService{}

type featureService struct{}

func (featureService) createDocumet(c *gin.Context) {
	var document featureModel
	err := c.BindJSON(&document)
	if err != nil {
		c.Error(errors.NewInvalidRequestError(err))
		c.Abort()
		return
	}
	result, err := repository.createDocument(c.Request.Context(), document)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result)
}

func (featureService) getDocumets(c *gin.Context) {
	filters := c.Query("filters")
	result, err := repository.getDocuments(c.Request.Context(), []byte(filters))
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result)
}
