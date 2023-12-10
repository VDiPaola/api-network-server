package helpers

import "net/http"

type RequestCallbackType func(response *http.Response, err error)
type ResponseCallbackType func(response *http.Response, err error)
