package region

import (
	"context"
	"mall/internal/dao"
	"mall/internal/model"
	"mall/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	sRegion struct{}
)

func init() {
	service.RegisterRegion(New())
}

func New() service.IRegion {
	return &sRegion{}
}

func (s *sRegion) List(ctx context.Context, parentId int) (res model.RegionList, err error) {
	err = dao.HiolabsRegion.Ctx(ctx).Where("parent_id = ?", parentId).Scan(&res)
	if err != nil {
		return nil, gerror.New("get region list error")
	}

	return
}
