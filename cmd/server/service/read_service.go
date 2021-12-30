package service

import (
	"context"
	pb "cqrs-grpc-test/api/contactpb"
	"cqrs-grpc-test/models"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReadService struct {
	pb.UnimplementedReadContactServiceServer
	Models *models.Models
}

func (svr *ReadService) Get(ctx context.Context, in *pb.GetContactReq) (*pb.GetContactResp, error) {

	c, err := svr.Models.GetContactById(ctx, in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%v", err))
	}

	return &pb.GetContactResp{Contact: c}, nil

}

func (svr *ReadService) GetList(ctx context.Context, in *pb.GetListContactReq) (*pb.GetListContactResp, error) {

	cc, total, err := svr.Models.GetContactList(ctx, in.Rang, in.Filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%v", err))
	}

	return &pb.GetListContactResp{
		Contact: cc,
		Total:   total,
		Rang:    in.Rang,
	}, nil

}
