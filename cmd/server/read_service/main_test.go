package main

import (
	"context"
	pb "cqrs-grpc-test/api/contactpb"
	utilpb "cqrs-grpc-test/api/utilpb"
	"fmt"
	"log"
	"testing"

	"google.golang.org/grpc"
)

func ListContact() (conn *grpc.ClientConn, c pb.ReadContactServiceClient, url string, err error) {
	addr := fmt.Sprintf("172.17.0.1:5001")

	conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	c = pb.NewReadContactServiceClient(conn)

	// md := metadata.New(map[string]string{"authorization": "Bearer +mytest2"})
	// ctx := metadata.NewOutgoingContext(context.Background(), md)

	resp, err := c.List(context.Background(), &pb.ListContactReq{
		Rang: &utilpb.QueryRange{
			PageSize: 10,
			Page:     0,
		},
		Filter: &pb.ListContactReq_Filter{},
	})
	fmt.Println(resp)

	if err != nil {
		return
	}

	return
}

func GetContact() (conn *grpc.ClientConn, c pb.ReadContactServiceClient, url string, err error) {
	addr := fmt.Sprintf("172.17.0.1:5001")

	conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	c = pb.NewReadContactServiceClient(conn)

	// md := metadata.New(map[string]string{"authorization": "Bearer +mytest2"})
	// ctx := metadata.NewOutgoingContext(context.Background(), md)

	resp, err := c.Get(context.Background(), &pb.GetContactReq{Uuid: "dee3b111-7288-4d9e-98bb-7f15da9a2e6f"})
	fmt.Println(resp)

	if err != nil {
		return
	}

	return
}

func TestListContact(t *testing.T) {
	conn, _, _, err := ListContact()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
}

func TestGetContact(t *testing.T) {
	conn, _, _, err := GetContact()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
}
