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

func (m *ContactModel) GetByUuid(ctx context.Context, id string) (*ent.Contact, error) {
	uid, err := util.UUidByStr(id)
	if err != nil {
		return nil, err
	}
	resp, err := m.Client.Contact.
		Query().
		Where(contact.UUID(uid)).
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

func (m *ContactModel) Update(ctx context.Context, id string, req *pb.ContactItem) error {
	uid, err := util.UUidByStr(id)
	if err != nil {
		return err
	}
	_, err = m.Client.Contact.
		Update().
		Where(contact.UUID(uid)).
		SetName(req.Name).
		SetPhone(req.Phone).
		SetUpdateTime(time.Now()).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (m *ContactModel) Delete(ctx context.Context, id string) error {
	uid, err := util.UUidByStr(id)
	if err != nil {
		return err
	}
	_, err = m.Client.Contact.
		Delete().
		Where(contact.UUID(uid)).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
