package service

import (
	"context"
	pb "cqrs-grpc-test/api/contactpb"
	"fmt"
	"log"
	"testing"

	"google.golang.org/grpc"
)

func CreateContact() (conn *grpc.ClientConn, c pb.WriteContactServiceClient, url string, err error) {
	addr := fmt.Sprintf("172.17.0.1:5002")

	conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	c = pb.NewWriteContactServiceClient(conn)

	// md := metadata.New(map[string]string{"authorization": "Bearer +mytest2"})
	// ctx := metadata.NewOutgoingContext(context.Background(), md)

	activity, err := c.Create(context.Background(), &pb.CreateContactReq{
		Item: &pb.ContactItem{
			Name:  "jason_test",
			Phone: "0912-123-456",
		},
	})
	fmt.Println(activity)
	if err != nil {
		return
	}

	return
}

func TestCreateContact(t *testing.T) {
	conn, _, _, err := CreateContact()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
}
