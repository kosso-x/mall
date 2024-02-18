package footprint

import (
	"context"
	"mall/internal/cmd"
	"mall/internal/dao"
	"mall/internal/model"
	"mall/internal/packed"
	"mall/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	sFootprint struct{}
)

func init() {
	service.RegisterFootprint(New())
}

func New() service.IFootprint {
	return &sFootprint{}
}

func (s *sFootprint) List(ctx context.Context, page, size int) (res *model.FootprintList, err error) {
	var (
		fpList  = model.FootprintList{}
		fpGoods []model.FootprintGoodsInfo
		count   int
	)

	g.
		Model("hiolabs_footprint hf").
		LeftJoin("hiolabs_goods hg", "hf.goods_id = hg.id").
		Fields("hf.id, hf.goods_id, hf.add_time, hg.name, hg.goods_brief, hg.retail_price, hg.list_pic_url, hg.goods_number, hg.min_cost_price").
		Where("hf.user_id = ?", cmd.CurrentUser.Id).
		Limit((page-1)*size, page*size).
		ScanAndCount(&fpGoods, &count, false)
	if err != nil {
		return
	}

	for _, goods := range fpGoods {
		fpList.Data = append(fpList.Data, model.FootprintInfoWithGoods{
			Id:      goods.Id,
			GoodsId: goods.GoodsId,
			AddTime: goods.AddTime,
			Goods: model.FootprintGoods{
				Name:           goods.Name,
				GoodsBrief:     goods.GoodsBrief,
				RetailPrice:    goods.RetailPrice,
				ListPicUrl:     goods.ListPicUrl,
				GoodsNumber:    goods.GoodsNumber,
				MinRetailPrice: goods.MinRetailPrice,
			},
		})
	}

	fpList.Count = count
	fpList.TotalPages = packed.TotalPages(ctx, fpList.Count, size)
	fpList.PageSize = size
	fpList.CurrentPage = page
	res = &fpList
	return
}

func (s *sFootprint) Delete(ctx context.Context, fpId int) (err error) {
	_, err = dao.HiolabsFootprint.Ctx(ctx).Delete("user_id = ? and id = ?", cmd.CurrentUser.Id, fpId)
	if err != nil {
		return gerror.New("delete footprint error")
	}

	return
}
