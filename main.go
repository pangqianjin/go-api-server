package main

import (
	"gitee.com/go-apiServer/service"
)

func main() {
	service := service.NewService()
	service.Start()
}
