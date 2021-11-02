package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/micro/micro/v3/service/logger"
	"user/domain/repository"
	service2 "user/domain/service"
	"user/handler"
	user "user/proto/user"

	"github.com/micro/go-micro/v2"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)

	srv.Init()

	db, err := gorm.Open("mysql", "")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)

	rp := repository.NewUserRepository(db)
	//rp.InitTable()

	userDataService := service2.NewUserDataService(rp)

	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})
	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
