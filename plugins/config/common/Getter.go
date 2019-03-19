package common

import (
	"github.com/goosmesh/goossidecar/utils/net"
)

var (
	DEFAULT_GOOS_ADDRESS = "http://goos:4321"
	API_PUB = "/api/pub"
	API_CONFIG = API_PUB + "/config/get"
)

func GetConfig(dataId string, groupId string, namespaceId string) (result string, err error) {
	params := make(map[string]string)
	if dataId != "" {
		params["dataId"] = dataId
	}
	if groupId != "" {
		params["groupId"] = groupId
	}
	if namespaceId != "" {
		params["namespaceId"] = namespaceId
	}
	return net.Get(DEFAULT_GOOS_ADDRESS + API_CONFIG, params)
}