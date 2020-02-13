package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"context"

	pb "github.com/bgoldovsky/shippy-client/proto/consignment"
	"github.com/micro/go-micro"
)

const (
	address         = "localhost:41249"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func nicePrint(c *pb.Consignment) {
	log.Println("===========CONSIGNMENT DATA===========")
	log.Printf("ID: %s\n", c.GetId())
	log.Printf("Description: %s\n", c.GetDescription())
	log.Printf("Weight: %d\n", c.GetWeight())
	log.Printf("Vessel: %s\n", c.GetVesselId())
	log.Println("======================================")
	log.Println()
}

func main() {
	service := micro.NewService(micro.Name("shippy-client"))
	service.Init()

	client := pb.NewShippingServiceClient("shippy-server", service.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		nicePrint(v)
	}
}