package helpers

import "net/http"

type RequestMethodType string

var RequestMethod = struct {
	GET  RequestMethodType
	POST RequestMethodType
}{
	GET:  http.MethodGet,
	POST: http.MethodPost,
}

func (r RequestMethodType) ToString() string {
	return string(r)
}
