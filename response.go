package hype

import (
	"net/http"
)

type Response struct {
	Actual *http.Response
	Data   []byte
	Error  error
}

func (response Response) Okay() bool {
	if response.Error == nil {
		return true
	}
	return false
}

func (response Response) Header(name string) string {
	return response.Actual.Header.Get(name)
}
