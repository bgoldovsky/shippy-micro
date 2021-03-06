package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/bgoldovsky/shippy-micro/shippyserver/proto/consignment"
	vesselProto "github.com/bgoldovsky/shippy-micro/shippyserver/proto/vessel"

	"github.com/micro/go-micro"
)

const (
	defaultHost = "localhost:27017"
)

func main() {

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	if err != nil {

		log.Fatalf("Could not connect to datastore with host %s - %v", host, err)
	}
	defer session.Close()

	srv := micro.NewService(

		micro.Name("shippyserver"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("shippyvessel", srv.Client())

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
