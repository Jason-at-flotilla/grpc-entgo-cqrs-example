package models

import (
	"context"
	pb "cqrs-grpc-test/api/contactpb"
	utilpb "cqrs-grpc-test/api/utilpb"
	"cqrs-grpc-test/ent"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Models) ContactEntToPb(data *ent.Contact) *pb.Contact {
	return &pb.Contact{
		Id: data.ID,
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

func (m *Models) UpdateContact(ctx context.Context, item *pb.ContactItem) error {

	_, err := m.ContactModel.Update(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func (m *Models) GetContactById(ctx context.Context, id int64) (*pb.Contact, error) {

	model, err := m.ContactModel.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return m.ContactEntToPb(model), nil
}

func (m *Models) GetContactList(ctx context.Context, r *utilpb.QueryRange, filter *pb.GetListContactReq_Filter) ([]*pb.Contact, int64, error) {

	if filter == nil {
		filter = &pb.GetListContactReq_Filter{}
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
