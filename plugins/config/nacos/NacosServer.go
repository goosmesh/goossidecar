package nacos

import (
	"github.com/goosmesh/goos/core/utils"
	"github.com/goosmesh/goos/core/utils/alg"
	"github.com/goosmesh/goossidecar/plugins/config/common"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var ServerPort = ":4322"

func StartServer()  {

	log.Info("start nacos proxy")

	//http.HandleFunc("/nacos/v1/cs/configs", nacosConfigs)

	modifyResponse := func(resp *http.Response, w http.ResponseWriter) error {
		if resp.StatusCode != 200 {
			return nil
		}
		rs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		bs := string(rs)
		if strings.Contains(bs, "code") {
			_, _ = w.Write([] byte(bs))
			return nil
		}
		okData, err := alg.RsaDecrypt(bs)
		if err != nil {
			return err
		}
		log.Info(okData)
		_, _ = w.Write([] byte(okData))
		return nil
	}
	configDirector := func(req *http.Request) {
		req.URL.RawQuery = strings.Replace(req.URL.RawQuery, "group=", "groupId=", 1)
	}
	listenerDirector := func(req *http.Request) {
		probeModify, err := utils.GetParameter("Listening-Configs", false, "", nil, req)
		if err != nil {
			return
		}
		nProbeModify := ParserMd5Data(probeModify)
		query := ""
		qs := strings.Split(req.URL.RawQuery, "&")
		for _, v := range qs {
			if strings.Index(v, "Listening-Configs=") == 0 {
				query += "Listening-Configs=" + url.QueryEscape(nProbeModify) + "&"
			} else {
				query += v + "&"
			}
		}
		log.Info(req.URL.RawQuery)
		req.URL.RawQuery = query[:len(query) - 1]
		log.Info(req.URL.RawQuery)
	}

	rpConfigs := common.RProxy{Host: common.DEFAULT_GOOS_HOST, Path: common.API_CONFIG, ModifyResponse: modifyResponse, Director:configDirector}
	rpConfigsListener := common.RProxy{Host: common.DEFAULT_GOOS_HOST, Path: common.API_CONFIG_LISTENER, Director:listenerDirector}
	http.HandleFunc("/nacos/v1/cs/configs", rpConfigs.ServeHTTP)
	http.HandleFunc("/nacos/v1/cs/configs/listener", rpConfigsListener.ServeHTTP)

	_ = http.ListenAndServe(ServerPort, nil)
}

//func nacosConfigs(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//
//	dataId, err := utils.GetParameter("dataId", false, "", w, r)
//	if err != nil {
//		return
//	}
//	groupId, err := utils.GetParameter("group", false, "", w, r)
//	if err != nil {
//		return
//	}
//	namespaceId, err := utils.GetParameter("tenant", true, "", w, r)
//	if err != nil {
//		return
//	}
//
//	config, _ := common.GetConfig(dataId, groupId, namespaceId)
//
//	_, _ = w.Write([]byte(config))
//
//}