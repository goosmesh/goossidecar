package common

import (
	"net/http"
	"net/http/httputil"
)

var (
	DEFAULT_GOOS_ADDRESS = "http://server.goos:4321"
	DEFAULT_GOOS_HOST = "server.goos:4321"
	API_PUB = "/api/pub"
	API_CONFIG = API_PUB + "/config/get"
	API_CONFIG_LISTENER = API_CONFIG + "/listener"
)

type RProxy struct {
	//  server.goos:4321
	Host string
	Path string
}

func (rproxy *RProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	director := func(req *http.Request) {
		req.URL.Scheme = r.URL.Scheme
		req.URL.Host = rproxy.Host
		req.URL.Path = rproxy.Path
	}

	proxy := httputil.ReverseProxy{Director:director}
	proxy.ServeHTTP(w, r)
}
