package net

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	Version = "goos-go-client:v0.0.1"
)

var httpClient = http.Client{
	Timeout: time.Duration(10000 * time.Millisecond),
}

func encodeUrl(urlString string, params map[string]string) string {

	if params == nil || len(params) == 0 {
		return urlString
	}
	if !strings.HasSuffix(urlString, "?") {
		urlString += "?"
	}

	u := url.Values{}
	for key, value := range params {
		u.Set(key, value)
	}

	return urlString + u.Encode()
}

func Get(url string, params map[string] string) (result string, err error) {

	if params == nil {
		params = make(map[string]string)
	}

	url = encodeUrl(url, params)

	req, err := http.NewRequest("GET",url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Client-Version", Version)
	response, err := httpClient.Do(req)

	if err != nil {
		return "", err
	}

	if response.StatusCode != 200 {
		return "", errors.New("error status : " + strconv.Itoa(response.StatusCode))
	}


	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	err = response.Body.Close()
	if err != nil {
		return "", err
	}

	bs := string(b)
	return bs, nil
}