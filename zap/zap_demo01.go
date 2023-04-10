package main

import (
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

func main() {
	InitLogger()
	defer logger.Sync()
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.baidu.com")
}

func simpleHttpGet(url string) {
	res, err := http.Get(url)
	if err != nil {

		logger.Error("Error fetching url...", zap.String("url", url), zap.Error(err))
	}else {
		logger.Info("Successfully...", zap.String("statu code", res.Status),zap.String("url", url))
		res.Body.Close()

	}
}

func InitLogger() {
	logger, _  = zap.NewProduction()
}
