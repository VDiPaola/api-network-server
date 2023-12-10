package options

import "github.com/VDiPaola/api-network-server/helpers"

var options helpers.OptionsType

func Set(newOptions helpers.OptionsType) {
	options = newOptions
}

func Get() helpers.OptionsType {
	return options
}
