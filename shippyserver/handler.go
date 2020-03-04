package main

import (
	"context"
	"log"

	pb "github.com/bgoldovsky/shippy-micro/shippyserver/proto/consignment"
	vesselProto "github.com/bgoldovsky/shippy-micro/shippyserver/proto/vessel"

	"gopkg.in/mgo.v2"
)

type service struct {
	session      *mgo.Session
	vesselClient vesselProto.VesselServiceClient
}

// GetRepo ...
func (s *service) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	defer s.GetRepo().Close()

	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err != nil {
		return err
	}

	log.Printf("Found vessel: %s (%s) \n", vesselResponse.Vessel.Name, vesselResponse.Vessel.Id)

	req.VesselId = vesselResponse.Vessel.Id

	err = s.GetRepo().Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	defer s.GetRepo().Close()
	consignments, err := s.GetRepo().GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
