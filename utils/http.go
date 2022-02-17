package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPReqConfig is the config struct for http requests.
type HTTPReqConfig struct {
	Method string
	URL    string
	Body   []byte
}

// NewHTTPClient creates a new http client.
func NewHTTPClient(timeout time.Duration) *http.Client {
	client := &http.Client{Timeout: timeout}
	return client
}

// HTTPRequest performs an HTTP request with config.
func HTTPRequest(client *http.Client, config HTTPReqConfig) ([]byte, error) {
	req, err := http.NewRequest(config.Method, config.URL, bytes.NewBuffer(config.Body))
	if err != nil {
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
