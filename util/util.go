package util

import (
	pbm "cqrs-grpc-test/api/utilpb"
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
