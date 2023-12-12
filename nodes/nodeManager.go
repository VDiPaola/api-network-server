package nodes

import (
	"time"

	"github.com/VDiPaola/api-network-server/models"
	"github.com/VDiPaola/api-network-server/options"
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
		AddNode(&node)
	}
}

func AddNode(node *models.Node) {
	node.IsActive = true
	node.LastResponseUnix = time.Now().Unix()
	node.RequestCount = 0
	node.PingArray = make([]int64, 0)
	node.ResponseTimeArray = make([]int64, 0)
	nodes = append(nodes, *node)
}

func cleanNodes() {
	//removes nodes that dont meet criteria
	options := options.Get()
	newNodes := []models.Node{}
	for _, node := range nodes {
		con1 := node.LastResponseUnix-time.Now().Unix() > int64(options.HealthCheckDuration.Seconds())
		con2 := node.IsActive
		con3 := node.Score+node.Priority > options.ScoreThreshold
		if con1 && con2 && con3 {
			newNodes = append(newNodes, node)
		}
	}

	nodes = newNodes
}
