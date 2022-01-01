package util

import (
	pbm "cqrs-grpc-test/api/utilpb"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

func GetOffset(r *pbm.QueryRange) int {
	offect := 0
	if r.Page == 0 {
		offect = 0
	} else {
		offect = int(r.PageSize * r.Page)
	}
	return offect
}

func UUidByStr(id string) (uuid.UUID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return uid, err
	}
	return uid, nil

}

func MarshalProto(message proto.Message) ([]byte, error) {
	value, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func UnmarshalProto(value string, data proto.Message) error {
	err := proto.Unmarshal([]byte(value), data)
	if err != nil {
		return err
	}
	return nil
}
