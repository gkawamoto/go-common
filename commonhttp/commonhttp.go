package commonhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// EncodingType type of encodings supported for the request payloads
type EncodingType int

const (
	// JSON specifies that the payload will be JSON
	JSON EncodingType = 1 << (10 * iota)
	// URLValues specifies that the payload will be url encoded
	URLValues
)

// NewClient creates a new HTTPClient
func NewClient() *HTTPClient {
	return &HTTPClient{
		client:       &http.Client{},
		EncodingType: JSON,
		Headers:      map[string]string{},
	}
}

// HTTPClient client wrapper for the http.Client
type HTTPClient struct {
	client       *http.Client
	EncodingType EncodingType
	Headers      map[string]string
}

// Get is a syntatic sugar for Request("GET" ...)
func (c *HTTPClient) Get(url string, params interface{}) ([]byte, int, error) {
	return c.Request("GET", url, params)
}

// Post is a syntatic sugar for Request("POST" ...)
func (c *HTTPClient) Post(url string, params interface{}) ([]byte, int, error) {
	return c.Request("POST", url, params)
}

// Delete is a syntatic sugar for Request("DELETE" ...)
func (c *HTTPClient) Delete(url string, params interface{}) ([]byte, int, error) {
	return c.Request("DELETE", url, params)
}

// Put is a syntatic sugar for Request("PUT" ...)
func (c *HTTPClient) Put(url string, params interface{}) ([]byte, int, error) {
	return c.Request("PUT", url, params)
}

// Request creates the request and executes it, returning the body, status code and eventual errors
func (c *HTTPClient) Request(method, url string, params interface{}) ([]byte, int, error) {
	var err error
	var data []byte
	var buffer bytes.Buffer
	data, err = encode(c.EncodingType, params)
	if err != nil {
		return nil, 0, err
	}
	_, err = buffer.Write(data)
	if err != nil {
		return nil, 0, err
	}
	var req *http.Request
	req, err = c.newRequestObject(method, url, &buffer)
	if err != nil {
		return nil, 0, err
	}
	var res *http.Response
	res, err = c.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, 0, err
	}
	return data, res.StatusCode, nil
}

func (c *HTTPClient) newRequestObject(method, url string, body io.Reader) (*http.Request, error) {
	var req *http.Request
	var err error
	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	var key, value string
	for key, value = range c.Headers {
		req.Header.Add(key, value)
	}
	return req, nil
}

func encode(encodingType EncodingType, params interface{}) ([]byte, error) {
	if params == nil {
		return []byte{}, nil
	}
	if encodingType == URLValues {
		if !isMap(params) {
			return nil, fmt.Errorf("invalid params (must be map[string]interface{})")
		}
		var paramsMap = params.(map[string]interface{})
		var key string
		var value interface{}
		var result url.Values
		for key, value = range paramsMap {
			result[key] = []string{value.(string)}
		}
		return []byte(result.Encode()), nil
	}
	if encodingType == JSON {
		return json.Marshal(params)
	}
	return nil, fmt.Errorf("invalid encodingType")
}

func isMap(data interface{}) bool {
	var ok bool
	_, ok = data.(map[string]interface{})
	return ok
}
