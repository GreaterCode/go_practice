/**
@author:admin
@date:2023/2/17
@note
*/

package main

import (
	"go.uber.org/zap"
	"net/http"
)

var sugaredlogger *zap.SugaredLogger

func main() {
	initLogger()
	defer logger.Sync()
	httpGet("www.google.com")
	httpGet("http://www.baidu.com")
}

func httpGet(url string) {
	sugaredlogger.Debugf("Trying to get http url %s", url)
	res, err := http.Get(url)
	if err != nil {
		 sugaredlogger.Errorf("Error getting http url %s: error:%v", url, err)
 	}else {
		 sugaredlogger.Infof("Successfully! status %s for URL %s", res.Status, url)
		 res.Body.Close()
 	}
}

func initLogger() {
	tmp_logger, _  := zap.NewProduction()
	sugaredlogger = tmp_logger.Sugar()
}
