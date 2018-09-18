package protoapigo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type protoapiClient struct {
	baseURL string
}

// ResponseHandler is the function to handle http response.
// It may bind result or return an error
type ResponseHandler func(res []byte, bizErr []byte, commonErr []byte) *Response

func (p *protoapiClient) SetBaseURL(url string) {
	p.baseURL = url
}

// CallAPI is used to call a API.
func (p *protoapiClient) CallAPI(req *Message, method string, URL string, handler ResponseHandler) *Response {
	client := &http.Client{}

	jsonStr, err := json.Marshal(req)
	if err != nil {
		return &Response{Err: err}
	}

	request, err := http.NewRequest(method, p.baseURL+URL, bytes.NewBuffer(jsonStr))
	if err != nil {
		return &Response{Err: err}
	}

	request.Header.Set("Content-Type", "application/json")
	rawResponse, err := client.Do(request)
	if err != nil {
		return &Response{Err: err}
	}
	defer rawResponse.Body.Close()

	body, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return &Response{Err: err}
	}

	switch rawResponse.StatusCode {
	case HAPPYPATH:
		return handler(body, nil, nil)
	case BIZERROR:
		return handler(nil, body, nil)
	case COMMONERROR:
		return handler(nil, nil, body)
	case ERROR:
		return &Response{Err: err}
	}
	return handler(nil, nil, nil)
}
