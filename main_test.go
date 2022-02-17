package main_test

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MichaelSimkin/go-template/config"
	"github.com/MichaelSimkin/go-template/server/errors"
	"github.com/MichaelSimkin/go-template/server/feature"
	"github.com/MichaelSimkin/go-template/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ = Describe("Feature", func() {
	httpClient := utils.NewHTTPClient(10 * time.Second)
	featureURL := "http://localhost:" + config.Service.Port + "/api/feature"
	var testObjectID primitive.ObjectID

	Context("Creating new document", func() {
		It("should create new document", func() {
			response, err := utils.HTTPRequest(httpClient, utils.HTTPReqConfig{
				Method: http.MethodPost,
				URL:    featureURL,
				Body:   []byte(`{"data":"Test data."}`),
			})
			Expect(err).To(BeNil())

			parsedRes := &feature.BaseModel{}
			err = json.Unmarshal(response, parsedRes)
			Expect(err).To(BeNil())
			Expect(parsedRes.ID).ToNot(BeEmpty())
			Expect(parsedRes.ID).To(BeAssignableToTypeOf(primitive.ObjectID{}))
			Expect(parsedRes.CreatedAt).To(BeTemporally("~", time.Now(), time.Second))
			Expect(parsedRes.UpdatedAt).To(BeTemporally("~", time.Now(), time.Second))
			Expect(parsedRes.CreatedAt).To(Equal(parsedRes.UpdatedAt))
			Expect(parsedRes.Data).To(Equal("Test data."))
			testObjectID = parsedRes.ID
		})

		It("should return error", func() {
			response, err := utils.HTTPRequest(httpClient, utils.HTTPReqConfig{
				Method: http.MethodPost,
				URL:    featureURL,
				Body:   []byte(`{"data":"Invalid data: $$$"}`),
			})
			Expect(err).To(BeNil())

			Expect(response).ToNot(BeNil())
			parsedErr := &errors.ServerError{}
			err = json.Unmarshal(response, parsedErr)
			Expect(err).To(BeNil())
			Expect(parsedErr.Code).To(Equal(http.StatusBadRequest))
			Expect(parsedErr.Message).To(Equal("Invalid request"))
		})
	})

	Context("Getting documents", func() {
		It("should get specific documents", func() {
			response, err := utils.HTTPRequest(httpClient, utils.HTTPReqConfig{
				Method: http.MethodGet,
				URL:    featureURL + `/?filters={"id":"` + testObjectID.Hex() + `"}`,
				Body:   []byte(`{}`),
			})
			Expect(err).To(BeNil())

			Expect(response).ToNot(BeEmpty())
			parsedRes := []*feature.BaseModel{}
			err = json.Unmarshal(response, &parsedRes)
			Expect(err).To(BeNil())
			Expect(len(parsedRes)).To(Equal(1))
			Expect(parsedRes[0].ID).To(Equal(testObjectID))
			Expect(parsedRes[0].Data).To(Equal("Test data."))
		})

		It("should get all documents", func() {
			response, err := utils.HTTPRequest(httpClient, utils.HTTPReqConfig{
				Method: http.MethodGet,
				URL:    featureURL + `/?filters={}`,
				Body:   []byte(`{}`),
			})
			Expect(err).To(BeNil())

			Expect(response).ToNot(BeEmpty())
			parsedRes := []*feature.BaseModel{}
			err = json.Unmarshal(response, &parsedRes)
			Expect(err).To(BeNil())
			Expect(len(parsedRes)).To(BeNumerically(">=", 1))
		})
	})
})
