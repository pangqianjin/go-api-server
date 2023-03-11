package main

import (
	"fmt"

	"gitee.com/go-apiServer/conf"
	"gitee.com/go-apiServer/libs"
	"gitee.com/go-apiServer/model"
	"gitee.com/go-apiServer/service"
)

func main() {
	property := model.Property{}
	libs.ReadProperty(conf.PropertyFileName, &property)

	fmt.Printf("%#v\n", property)
	service := service.NewService(property)
	service.Start()
}
