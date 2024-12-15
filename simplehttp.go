package simplehttp

import (
	"bytes"
	"io"
	"net/http"
)

type HttpReqest struct {
	method  string
	url     string
	request http.Request
	body    io.Reader
}

var globalClient = &http.Client{}

func CustomClient(client *http.Client) {
	if client != nil {
		globalClient = client
	}
}

func CreateHttpRequest(url string) HttpReqest {
	return HttpReqest{
		method: http.MethodGet,
		body:   nil,
		url:    url,
	}
}

func (h *HttpReqest) Url(u string) *HttpReqest {
	h.url = u

	return h
}

func (h *HttpReqest) Get() *HttpReqest {
	return h.SetMethod(http.MethodGet)
}

func (h *HttpReqest) Post() *HttpReqest {
	return h.SetMethod(http.MethodPost)
}

func (h *HttpReqest) Put() *HttpReqest {
	return h.SetMethod(http.MethodPut)
}

func (h *HttpReqest) Delete() *HttpReqest {
	return h.SetMethod(http.MethodDelete)
}

func (h *HttpReqest) SetMethod(method string) *HttpReqest {
	h.method = method

	return h
}

func (h *HttpReqest) Body(body []byte) *HttpReqest {
	h.body = bytes.NewReader(body)
	return h
}

func (h *HttpReqest) AddHeader(key, value string) *HttpReqest {
	if h.request.Header == nil {
		h.request.Header = make(http.Header)
	}

	h.request.Header.Add(key, value)

	return h
}

func (h *HttpReqest) Execute() (*http.Response, error) {
	return h.ExecuteWithClient(*globalClient)
}

func (h *HttpReqest) ExecuteWithClient(client http.Client) (*http.Response, error) {
	req, err := http.NewRequest(
		h.method,
		h.url,
		h.body,
	)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil

}
