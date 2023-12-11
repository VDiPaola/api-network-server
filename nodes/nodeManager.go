package nodes

import (
	"time"

	"github.com/VDiPaola/api-network-server/models"
)

var nodes = make([]models.Node, 0)

func GetNodes() []models.Node {
	return nodes
}

func GetNode(parsedNode models.Node) *models.Node {
	for _, node := range nodes {
		if node.IP == parsedNode.IP {
			return &node
		}
	}

	return nil
}

func NodeExists(parsedNode models.Node) bool {
	for _, node := range nodes {
		if node.IP == parsedNode.IP {
			return true
		}
	}

	return false
}

func AddNodes(nodeArray []models.Node) {
	for _, node := range nodeArray {
		AddNode(node)
	}
}

func AddNode(node models.Node) {
	node.IsActive = true
	node.LastResponseUnix = time.Now().Unix()
	node.RequestCount = 0
	nodes = append(nodes, node)
}
