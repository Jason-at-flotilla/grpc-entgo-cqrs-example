package server

import (
	"context"
	pb "cqrs-grpc-test/api/contactpb"

	google_protobuf "github.com/golang/protobuf/ptypes/empty"

	"cqrs-grpc-test/models"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WriteService struct {
	pb.UnimplementedWriteContactServiceServer
	Models *models.Models
}

func (svr *WriteService) Create(ctx context.Context, in *pb.CreateContactReq) (*pb.CreateContactResp, error) {

	c, err := svr.Models.CreateContact(ctx, in.Item)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%v", err))
	}

	return &pb.CreateContactResp{Contact: c}, nil
}

func (svr *WriteService) Update(ctx context.Context, in *pb.UpdateContactReq) (*pb.UpdateContactResp, error) {

	err := svr.Models.UpdateContact(ctx, in.Item)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%v", err))
	}

	return &pb.UpdateContactResp{Item: in.Item}, nil

}

func (svr *WriteService) Delete(ctx context.Context, in *pb.DeleteContactReq) (*google_protobuf.Empty, error) {

	err := svr.Models.ContactModel.Delete(ctx, in.Uuid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%v", err))
	}

	return &google_protobuf.Empty{}, nil

}
