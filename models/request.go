package models

import (
	"github.com/VDiPaola/api-network-server/helpers"
)

type Request struct {
	Endpoint string
	Body     interface{}
	Method   helpers.RequestMethodType
}
