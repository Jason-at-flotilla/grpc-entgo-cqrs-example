package entclient

import (
	"context"
	"cqrs-grpc-test/ent"
	"fmt"
	"sync"
)

type MyEnt struct {
	Ent *ent.Client
}

var client *MyEnt
var once sync.Once
var err error
var dataSource = ""
var db = ""

func Init(EnvDataSource string, EnvDb string) {
	dataSource = EnvDataSource
	db = EnvDb
}

func GetInstance() (*MyEnt, error) {
	once.Do(func() {
		client = &MyEnt{}
		client.Ent, err = ent.Open(db, dataSource)
	})
	return client, err
}
func (m *MyEnt) GetTx(ctx context.Context) (*ent.Tx, error) {
	tx, err := client.Ent.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (m *MyEnt) Rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
