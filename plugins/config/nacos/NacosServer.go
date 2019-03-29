package nacos

import (
	"github.com/goosmesh/goossidecar/plugins/config/common"
	"github.com/prometheus/common/log"
	"net/http"
)

func StartServer()  {

	log.Info("start nacos proxy")

	//http.HandleFunc("/nacos/v1/cs/configs", nacosConfigs)
	rpConfigs := common.RProxy{Host: common.DEFAULT_GOOS_HOST, Path: common.API_CONFIG}
	rpConfigsListener := common.RProxy{Host: common.DEFAULT_GOOS_HOST, Path: common.API_CONFIG_LISTENER}
	http.HandleFunc("/nacos/v1/cs/configs", rpConfigs.ServeHTTP)
	http.HandleFunc("/nacos/v1/cs/configs/listener", rpConfigsListener.ServeHTTP)

	_ = http.ListenAndServe(":4323", nil)
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