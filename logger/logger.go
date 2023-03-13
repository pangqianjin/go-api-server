package logger

import (
	"fmt"

	"gitee.com/go-apiServer/conf"
	"gitee.com/go-apiServer/libs"
	"gitee.com/go-apiServer/model"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/logger"
)

func CustomLogger() (customLogger context.Handler) {
	config := model.LoggerConfig{}
	libs.ReadFileToModel(conf.LoggerConfigFileName, &config)
	fmt.Printf("%#v\n", config)

	customLogger = logger.New(logger.Config{
		Status: config.Status,
		IP:     config.IP,
		Method: config.Method,
		Path:   config.Path,
		Query:  config.Query,
	})
	return
}
