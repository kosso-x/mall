package search

import (
	"context"

	v1 "mall/api/search/v1"
	"mall/internal/service"
)

func (c *ControllerV1) SearchIndex(ctx context.Context, req *v1.SearchIndexReq) (res *v1.SearchIndexRes, err error) {
	r, err := service.Search().Index(ctx)
	if err != nil {
		return
	}

	res = &v1.SearchIndexRes{
		SearchIndex: r,
	}
	return
}
