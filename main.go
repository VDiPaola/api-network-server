package main

import (
	"github.com/VDiPaola/api-network-server/helpers"
	"github.com/VDiPaola/api-network-server/models"
	"github.com/VDiPaola/api-network-server/options"
)

var nodes = make([]models.Node, 0)

//request queue

func Init(newOptions helpers.OptionsType) {
	//set options
	options.Set(newOptions)

	//caching

}

//sort based on score
// sort.Slice(readyNodes, func(i, j int) bool {
// 	return readyNodes[i].Score < readyNodes[j].Score
// })
