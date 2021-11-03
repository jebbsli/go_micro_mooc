package main

import (
	"category/common"
	"category/domain/repository"
	service2 "category/domain/service"
	"github.com/jinzhu/gorm"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"category/handler"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-plugins/registry/consul/v2"

	category "category/proto/category"
)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Fatal(err)
	}

	// 注册中心
	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8082"),
		micro.Registry(consulRegister),
	)

	// 获取mysql配置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig)
	db, err := gorm.Open("mysql", mysqlInfo.Host + ":" + mysqlInfo.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SingularTable(true)

	// Initialise service
	service.Init()

	categoryDataService := service2.NewCategoryDataService(repository.NewCategoryRepository(db))


	// Register Handler
	category.RegisterCategoryHandler(service.Server(), &handler.Category{CategoryDataService: categoryDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
