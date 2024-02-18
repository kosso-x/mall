// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package search

import (
	"context"

	"mall/api/search/v1"
)

type ISearchV1 interface {
	SearchIndex(ctx context.Context, req *v1.SearchIndexReq) (res *v1.SearchIndexRes, err error)
	SearchHelper(ctx context.Context, req *v1.SearchHelperReq) (res *v1.SearchHelperRes, err error)
	SearchClearHistory(ctx context.Context, req *v1.SearchClearHistoryReq) (res *v1.SearchClearHistoryRes, err error)
}
