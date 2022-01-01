package models

import (
	"context"
	pb "cqrs-grpc-test/api/contactpb"
	utilpb "cqrs-grpc-test/api/utilpb"
	"cqrs-grpc-test/ent"
	"cqrs-grpc-test/pkg/redis"
	"cqrs-grpc-test/util"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Models) ContactEntToPb(data *ent.Contact) *pb.Contact {
	return &pb.Contact{
		Uuid: data.UUID.String(),
		Item: &pb.ContactItem{
			Name:  data.Name,
			Phone: data.Phone,
		},
		CreatedAt: timestamppb.New(data.CreateTime),
		UpdatedAt: timestamppb.New(data.UpdateTime),
	}

}

func (m *Models) CreateContact(ctx context.Context, item *pb.ContactItem) (*pb.Contact, error) {

	model, err := m.ContactModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return m.ContactEntToPb(model), nil
}

func (m *Models) UpdateContact(ctx context.Context, item *pb.ContactItem, uuid string) error {

	r := redis.GetClient(int(redis.DB_CACHE))
	r.Del(fmt.Sprintf("/contact/%s", uuid))

	err := m.ContactModel.Update(ctx, uuid, item)
	if err != nil {
		return err
	}

	return nil
}

func (m *Models) GetContactById(ctx context.Context, uuid string) (*pb.Contact, error) {
	item := &pb.Contact{}
	r := redis.GetClient(int(redis.DB_CACHE))
	v, err := r.Get(fmt.Sprintf("/contact/%s", uuid))
	if err != nil {
		model, err := m.ContactModel.GetByUuid(ctx, uuid)
		if err != nil {
			return nil, err
		}
		item = m.ContactEntToPb(model)
		setValue, err := util.MarshalProto(item)
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("%v", err))
		}
		r.Set(fmt.Sprintf("/contact/%s", uuid), setValue, 0)
		return item, nil
	}
	util.UnmarshalProto(v, item)

	return item, nil
}

func (m *Models) GetContactList(ctx context.Context, r *utilpb.QueryRange, filter *pb.ListContactReq_Filter) ([]*pb.Contact, int64, error) {

	if filter == nil {
		filter = &pb.ListContactReq_Filter{}
	}

	model, err := m.ContactModel.GetByFilter(ctx, filter.Name, r)
	if err != nil {
		return nil, 0, err
	}

	total, err := m.ContactModel.GetByFilterCount(ctx, filter.Name)
	if err != nil {
		return nil, 0, err
	}

	cc := []*pb.Contact{}
	for i := 0; i < len(model); i++ {
		cc = append(cc, m.ContactEntToPb(model[i]))
	}

	return cc, int64(total), nil
}

func (m *Models) DeleteContact(ctx context.Context, uuid string) error {

	r := redis.GetClient(int(redis.DB_CACHE))
	r.Del(fmt.Sprintf("/contact/%s", uuid))

	err := m.ContactModel.Delete(ctx, uuid)
	if err != nil {
		return err
	}

	return nil

}
