package models

import "cqrs-grpc-test/models/postgres"

type Models struct {
	ContactModel *postgres.ContactModel
}
