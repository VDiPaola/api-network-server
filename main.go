package main

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/VDiPaola/api-network-server/helpers"
	"github.com/VDiPaola/api-network-server/models"
)

type Options struct {
	RateLimitAmount   uint //amount of request per RateLimitInterval
	RateLimitInterval time.Duration
	//MaxRequestsPerMinute uint
	MaxResponseDuration time.Duration //max time a node is allowed to respond in
	NodeRedundencyCount uint          //how many nodes to ask a request to simultaneously
	HealthCheckDuration time.Duration //how long to wait before considering a node inactive
	CacheUpdateInterval time.Duration
	HasCache            bool
	ScoreThreshold      int64
}

var options = Options{
	RateLimitAmount:     300,
	RateLimitInterval:   time.Duration(time.Minute * 5),
	MaxResponseDuration: time.Duration(time.Millisecond * 100),
	NodeRedundencyCount: 2,
	HealthCheckDuration: time.Duration(time.Minute * 15),
	CacheUpdateInterval: time.Duration(time.Minute * 5),
	HasCache:            true,
	ScoreThreshold:      10,
}

var nodes = make([]models.Node, 0)

//request queue

func Init(newOptions Options) {
	//load options
	options = newOptions

	//caching

}

func AddNodes(nodeArray []models.Node) {
	for _, node := range nodeArray {
		AddNode(node)
	}
}

func AddNode(node models.Node) {
	nodes = append(nodes, node)
}

func healthCheck(node models.Node) {
	//pinged from the node to show its active
}

func selectNodes() []models.Node {
	//returns array of nodes selected for requests
	if len(nodes) == int(options.NodeRedundencyCount) {
		return nodes
	}

	//select nodes from best performing
	readyNodes := getReadyNodes()

	var selectedNodes = make([]models.Node, 0)

	for i := 0; i < int(options.NodeRedundencyCount); i++ {
		//add selected node them remove from readyNodes
		selectedIndex := rand.Intn(len(readyNodes))
		selectedNodes = append(selectedNodes, readyNodes[selectedIndex])
		readyNodes = helpers.Remove(readyNodes, selectedIndex)
	}

	return selectedNodes
}

func getReadyNodes() []models.Node {
	//get nodes that are valid to make a request
	newNodes := []models.Node{}
	for i := range nodes {
		requestCountCheck(nodes[i])
		if nodes[i].RequestCount < options.RateLimitAmount {
			newNodes = append(newNodes, nodes[i])
		}
	}

	return newNodes
}

func requestCountCheck(node models.Node) {
	//resets request count if RateLimitInterval passed
	if node.NextRequestCountResetTime.Unix() <= time.Now().Unix() {
		node.RequestCount = 0
		node.NextRequestCountResetTime = time.Now().Add(options.RateLimitInterval)
	}
}

func cleanNodes() {
	//removes nodes that dont meet criteria
	newNodes := []models.Node{}
	for i := range nodes {
		con1 := nodes[i].LastResponseUnix-time.Now().Unix() > int64(options.HealthCheckDuration.Seconds())
		con2 := nodes[i].IsActive
		con3 := nodes[i].Score+nodes[i].Priority > options.ScoreThreshold
		if con1 && con2 && con3 {
			newNodes = append(newNodes, nodes[i])
		}
	}

	nodes = newNodes
}

func Request(endpoint string, method helpers.RequestMethodType, body interface{}, callback helpers.ResponseCallbackType) {
	//get nodes to run requests on
	hasSent := false
	selectedNodes := selectNodes()
	for _, node := range selectedNodes {
		//run requests
		processRequest(node, endpoint, method, body, func(response *http.Response, err error) {
			if !hasSent {
				hasSent = true
				callback(response, err)
			}
		})
	}
}

func processRequest(node models.Node, endpoint string, method helpers.RequestMethodType, body interface{}, callback helpers.RequestCallbackType) {
	//run request on node

	//time request
	start := time.Now()

	//serialise body
	jsonBody, err := json.Marshal(models.Request{
		Endpoint: endpoint,
		Body:     body,
		Method:   method,
	})

	jsonBuffer := bytes.NewBuffer(jsonBody)

	if err != nil {
		callback(nil, err)
		return
	}

	req, err := http.NewRequest(method.ToString(), endpoint, jsonBuffer)

	if err != nil {
		callback(nil, err)
		return
	}

	res, err := http.DefaultClient.Do(req)

	diff := time.Since(start)

	//check valid response
	if res.StatusCode != 200 {
		callback(nil, err)
		return
	}

	//check times
	if diff > options.MaxResponseDuration {
		node.Score -= 1
	} else if diff < time.Millisecond*20 {
		node.Score += 1
	}

	//return results in callback function
	callback(res, nil)
}

//sort based on score
// sort.Slice(readyNodes, func(i, j int) bool {
// 	return readyNodes[i].Score < readyNodes[j].Score
// })
