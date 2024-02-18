package search

import (
	"context"

	v1 "mall/api/search/v1"
	"mall/internal/service"
)

func (c *ControllerV1) SearchHelper(ctx context.Context, req *v1.SearchHelperReq) (res *v1.SearchHelperRes, err error) {
	r, err := service.Search().Helper(ctx, req.Keyword)
	if err != nil {
		return
	}
	res = &v1.SearchHelperRes{
		HelpKeywords: r,
	}

	return
}
