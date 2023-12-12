package nodes

import (
	"time"

	"github.com/VDiPaola/api-network-server/models"
)

func HealthCheck(node models.Node) models.Node {
	//pinged from the node to show its active

	//create node if not exists
	currentNode := GetNode(node)
	if currentNode == nil {
		node.Priority = 0
		node.Score = 100
		AddNode(&node)
	} else {
		//update time
		currentNode.LastResponseUnix = time.Now().Unix()
		currentNode.IsActive = true
		node = *currentNode
	}

	return node
}
