package service

import (
	"context"
	pb "cqrs-grpc-test/api/contactpb"
	utilpb "cqrs-grpc-test/api/utilpb"
	"fmt"
	"log"
	"testing"

	"google.golang.org/grpc"
)

func GetListContact() (conn *grpc.ClientConn, c pb.ReadContactServiceClient, url string, err error) {
	addr := fmt.Sprintf("172.17.0.1:5002")

	conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	c = pb.NewReadContactServiceClient(conn)

	// md := metadata.New(map[string]string{"authorization": "Bearer +mytest2"})
	// ctx := metadata.NewOutgoingContext(context.Background(), md)

	resp, err := c.GetList(context.Background(), &pb.GetListContactReq{
		Rang: &utilpb.QueryRange{
			PageSize: 10,
			Page:     0,
		},
		Filter: &pb.GetListContactReq_Filter{},
	})
	fmt.Println(resp.Total)

	if err != nil {
		return
	}

	return
}

func TestGetContact(t *testing.T) {
	conn, _, _, err := GetListContact()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
}
