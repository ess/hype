package hype

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	actual *http.Request
	raw    *http.Client
	err    error
}

func (request *Request) okay() bool {
	return request.err == nil
}

func (request *Request) WithHeader(header *Header) *Request {
	return request.WithHeaderSet(header)
}

func (request *Request) WithHeaderSet(headers ...*Header) *Request {
	if request.okay() {
		for _, header := range headers {
			request.actual.Header.Add(header.Name, header.Value)
		}
	}

	return request
}

func (request *Request) Response() Response {
	if !request.okay() {
		return Response{nil, nil, request.err}
	}

	response, err := request.raw.Do(request.actual)
	if err != nil {
		return Response{nil, nil, err}
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Response{nil, nil, err}
	}

	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return Response{nil, nil, fmt.Errorf("response status: %d", response.StatusCode)}
	}

	return Response{response, body, nil}
}
