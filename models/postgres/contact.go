package postgres

import (
	"context"
	pb "cqrs-grpc-test/api/contactpb"
	pbm "cqrs-grpc-test/api/utilpb"
	"cqrs-grpc-test/ent"
	"cqrs-grpc-test/ent/contact"
	"cqrs-grpc-test/util"
	"time"

	"entgo.io/ent/dialect/sql"
)

type ContactModel struct {
	Client *ent.Client
}

func (m *ContactModel) Create(ctx context.Context, req *pb.ContactItem) (*ent.Contact, error) {

	resp, err := m.Client.Contact.
		Create().
		SetName(req.Name).
		SetPhone(req.Phone).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ContactModel) GetById(ctx context.Context, id int64) (*ent.Contact, error) {

	resp, err := m.Client.Contact.
		Query().
		Where(contact.ID(id)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ContactModel) GetByFilter(ctx context.Context, name string, r *pbm.QueryRange) ([]*ent.Contact, error) {

	resp, err := m.Client.Contact.
		Query().
		Where(
			contact.And(
				func(s *sql.Selector) {
					if name != "" {
						s.Where(sql.Like(contact.FieldName, "%"+name+"%"))
					}
				},
			),
		).
		Limit(int(r.PageSize)).
		Offset(util.GetOffset(r)).
		Order(ent.Asc(contact.FieldCreateTime)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ContactModel) GetByFilterCount(ctx context.Context, name string) (int, error) {

	count, err := m.Client.Contact.
		Query().
		Where(
			contact.And(
				func(s *sql.Selector) {
					if name != "" {
						s.Where(sql.Like(contact.FieldName, "%"+name+"%"))
					}
				},
			),
		).
		Count(ctx)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (m *ContactModel) Update(ctx context.Context, req *pb.ContactItem) (int, error) {

	id, err := m.Client.Contact.
		Update().
		SetName(req.Name).
		SetPhone(req.Phone).
		SetUpdateTime(time.Now()).
		Save(ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *ContactModel) Delete(ctx context.Context, id int64) error {

	err := m.Client.Contact.
		DeleteOneID(id).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
