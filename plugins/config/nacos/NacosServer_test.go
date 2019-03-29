package nacos

import (
	"github.com/goosmesh/goossidecar/plugins/config/common"
	"testing"
)

func TestStartServer(t *testing.T) {
	common.DEFAULT_GOOS_HOST = "server.goos:3321"
	ServerPort = ":3322"
	StartServer()
}