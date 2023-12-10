package caching

import "github.com/VDiPaola/api-network-server/models"

type ResponseCache interface {
	Set(key string, value *models.Node)
	Get(key string) *models.Node
}
