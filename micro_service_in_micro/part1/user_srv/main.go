package main

import (
	"fmt"

	"github.com/pylrichard/micro_study/micro_service_in_micro/part1/user_srv/basic"
	"github.com/pylrichard/micro_study/micro_service_in_micro/part1/user_srv/basic/config"
	"github.com/pylrichard/micro_study/micro_service_in_micro/part1/user_srv/handler"
	"github.com/pylrichard/micro_study/micro_service_in_micro/part1/user_srv/model"
	user "github.com/pylrichard/micro_study/micro_service_in_micro/part1/user_srv/proto/user"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"
)

func main() {
	//初始化配置、数据库等
	basic.Init()

	//使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
		micro.Version("latest"),
	)

	service.Init()

	user.RegisterUserHandler(service.Server(), new(handler.User))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdConfig = config.GetEtcdConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string {
		fmt.Sprintf("%s:%d", etcdConfig.GetHost(), etcdConfig.GetPort())
	}
}