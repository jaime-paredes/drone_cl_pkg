package helpers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

//DefaultClient initial setup
var DefaultClient = &Client{HTTPClient: http.DefaultClient}

//Client of http
type Client struct {
	HTTPClient *http.Client
}

//Request basic struc of a request petition
type Request struct {
	Method  string
	BaseURL string
	Headers map[string]string
	Body    []byte
}

//Response default struct
type Response struct {
	StatusCode int
	Body       []byte
	Headers    map[string][]string
}

//APIControl main function that will control the flow
func (c *Client) APIControl(request Request) (*Response, error) {
	req, err := BuildRequest(request)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	return BuildResponse(res)

}

//BuildRequest will form the struct to do the request
func BuildRequest(request Request) (*http.Request, error) {

	req, err := http.NewRequest(strings.ToUpper(request.Method), request.BaseURL, bytes.NewBuffer(request.Body))
	if err != nil {
		return req, err
	}
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")
	return req, err
}

//BuildResponse will form the struct after the call to some api
func BuildResponse(res *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(res.Body)
	response := Response{
		StatusCode: res.StatusCode,
		Body:       body,
		Headers:    res.Header,
	}
	res.Body.Close()
	return &response, err
}

//APICall func that execute APICall
func APICall(request Request) (*Response, error) {
	return DefaultClient.APIControl(request)
}
