package net

import (
	"bytes"
	"github.com/goosmesh/goos/core/utils/alg"
	"github.com/pkg/errors"
	"io"
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

	// 解码数据
	return alg.RsaDecrypt(bs)
}

func Post(url string, params map[string] string) (result string, err error)  {

	if params == nil {
		params = make(map[string]string)
	}

	url = encodeUrl(url, params)

	req, err := http.NewRequest("POST",url, nil)
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

	// 解码数据
	return alg.RsaDecrypt(bs)
}

func WriteBody(body io.Reader, req *http.Request) {
	if body != nil {
		switch v := body.(type) {
		case *bytes.Buffer:
			req.ContentLength = int64(v.Len())
			buf := v.Bytes()
			req.GetBody = func() (io.ReadCloser, error) {
				r := bytes.NewReader(buf)
				return ioutil.NopCloser(r), nil
			}
		case *bytes.Reader:
			req.ContentLength = int64(v.Len())
			snapshot := *v
			req.GetBody = func() (io.ReadCloser, error) {
				r := snapshot
				return ioutil.NopCloser(&r), nil
			}
		case *strings.Reader:
			req.ContentLength = int64(v.Len())
			snapshot := *v
			req.GetBody = func() (io.ReadCloser, error) {
				r := snapshot
				return ioutil.NopCloser(&r), nil
			}
		default:
			// This is where we'd set it to -1 (at least
			// if body != NoBody) to mean unknown, but
			// that broke people during the Go 1.8 testing
			// period. People depend on it being 0 I
			// guess. Maybe retry later. See Issue 18117.
		}
		// For client requests, Request.ContentLength of 0
		// means either actually 0, or unknown. The only way
		// to explicitly say that the ContentLength is zero is
		// to set the Body to nil. But turns out too much code
		// depends on NewRequest returning a non-nil Body,
		// so we use a well-known ReadCloser variable instead
		// and have the http package also treat that sentinel
		// variable to mean explicitly zero.
		if req.GetBody != nil && req.ContentLength == 0 {
			req.Body = http.NoBody
			req.GetBody = func() (io.ReadCloser, error) { return http.NoBody, nil }
		}
	}
}