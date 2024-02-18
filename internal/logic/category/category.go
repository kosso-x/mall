package category

import (
	"context"
	"mall/internal/dao"
	"mall/internal/model"
	"mall/internal/model/do"
	"mall/internal/packed"
	"mall/internal/service"

	"github.com/gogf/gf/errors/gerror"
)

type (
	sCategory struct{}
)

func init() {
	service.RegisterCategory(New())
}

func New() service.ICategory {
	return &sCategory{}
}

func (s *sCategory) List(ctx context.Context) (list model.CategoryList, err error) {
	err = dao.
		HiolabsCategory.
		Ctx(ctx).
		Order("sort_order asc").
		Scan(&list)

	if err != nil {
		return nil, gerror.New("get categories error")
	}

	return
}

func (s *sCategory) One(ctx context.Context, id int) (current model.CurrentCategory, err error) {
	dao.
		HiolabsCategory.
		Ctx(ctx).
		Fields("id", "name", "img_url", "p_height").
		Where(do.HiolabsCategory{
			Id: id,
		}).Scan(&current)

	return
}

func (s *sCategory) CurrentList(ctx context.Context, id, page, size int) (res model.CategoryCurrentList, err error) {
	var (
		goods []model.ShortGoods
		count int
	)

	if id == 0 {
		err = dao.
			HiolabsGoods.
			Ctx(ctx).
			Fields("id", "name", "goods_brief", "min_retail_price", "list_pic_url", "goods_number").
			Where("is_on_sale = ? and is_delete = ?", 1, 0).
			Limit((page-1)*size, page*size).
			Order("sort_order asc").
			ScanAndCount(&goods, &count, false)

		if err != nil {
			return res, gerror.New("get all-category-goods-count error")
		}
	} else {
		err = dao.
			HiolabsGoods.
			Ctx(ctx).
			Fields("id", "name", "goods_brief", "min_retail_price", "list_pic_url", "goods_number").
			Where("is_on_sale = ? and is_delete = ? and category_id = ?", 1, 0, id).
			Limit((page-1)*size, page*size).
			Order("sort_order asc").
			ScanAndCount(&goods, &count, false)

		if err != nil {
			return res, gerror.New("get category-goods-count error")
		}
	}

	res = model.CategoryCurrentList{
		Count:       count,
		TotalPages:  packed.TotalPages(ctx, count, size),
		PageSize:    size,
		CurrentPage: page,
		Data:        goods,
	}

	return
}
