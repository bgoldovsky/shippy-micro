package main

import (
	"fmt"
	"log"

	pb "github.com/bgoldovsky/shippy-micro/userserver/proto/user"
	"github.com/micro/go-micro"
)

func main() {
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("could not connect to database postgres. %v", err)
	}

	db.AutoMigrate(&pb.User{})
	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("userserver"),
		micro.Version("latest"),
	)
	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
