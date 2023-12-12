package nodes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/VDiPaola/api-network-server/helpers"
	"github.com/VDiPaola/api-network-server/models"
	"github.com/VDiPaola/api-network-server/options"
)

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
	options := options.Get()
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

	//create request
	req, err := http.NewRequest(helpers.RequestMethod.POST.ToString(), fmt.Sprintf("http://%v:%v/api-network/request", node.IP, node.Port), jsonBuffer)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		callback(nil, err)
		return
	}

	//send request
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

func selectNodes() []models.Node {
	//returns array of nodes selected for requests
	options := options.Get()
	if len(nodes) == int(options.NodeRedundencyCount) {
		return nodes
	}

	//select nodes from best performing
	readyNodes := getReadyNodes()

	var selectedNodes = make([]models.Node, 0)

	for i := 0; i < int(options.NodeRedundencyCount); i++ {
		//break if no nodes left
		if len(readyNodes) <= 0 {
			break
		}
		//add selected node them remove from readyNodes
		selectedIndex := rand.Intn(len(readyNodes))
		selectedNodes = append(selectedNodes, readyNodes[selectedIndex])
		readyNodes = helpers.Remove(readyNodes, selectedIndex)
	}

	return selectedNodes
}

func getReadyNodes() []models.Node {
	//get nodes that are valid to make a request
	options := options.Get()
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
	options := options.Get()
	//resets request count if RateLimitInterval passed
	if node.NextRequestCountResetTime.Unix() <= time.Now().Unix() {
		node.RequestCount = 0
		node.NextRequestCountResetTime = time.Now().Add(options.RateLimitInterval)
	}
}
