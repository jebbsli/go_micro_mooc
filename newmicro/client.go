package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	imooc "newmicro/proto/cap"
)

func main() {
	// 创建新的服务
	service := micro.NewService(
		micro.Name("cap.imooc.client"),
		)

	// 初始化
	service.Init()

	capImooc := imooc.NewCapService("cap.imooc.server", service.Client())

	res, err := capImooc.SayHello(context.TODO(), &imooc.SayRequest{Message: "什么口号？"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Answer)
}
