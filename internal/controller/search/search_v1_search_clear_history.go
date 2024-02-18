package search

import (
	"context"

	v1 "mall/api/search/v1"
	"mall/internal/service"
)

func (c *ControllerV1) SearchClearHistory(ctx context.Context, req *v1.SearchClearHistoryReq) (res *v1.SearchClearHistoryRes, err error) {
	var is_delete = 1
	err = service.Search().ClearHistory(ctx)
	if err != nil {
		is_delete = 0
		return
	}

	res = &v1.SearchClearHistoryRes{
		IsDelete: is_delete,
	}

	return
}
