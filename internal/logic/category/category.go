package category

import (
	"context"
	"mall/internal/dao"
	"mall/internal/model"
	"mall/internal/model/do"
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

func (s *sCategory) CurrentList(ctx context.Context, Id, Page, Size int) (res model.CategoryCurrentList, err error) {
	var goods []model.ShortGoods
	count, err := dao.
		HiolabsGoods.
		Ctx(ctx).
		Where("is_on_sale = ? and is_delete = ? and category_id = ?", 1, 0, Id).
		Count()

	if err != nil {
		return res, gerror.New("get category-goods-count error")
	}

	dao.
		HiolabsGoods.
		Ctx(ctx).
		Fields("name", "id", "goods_brief", "min_retail_price", "list_pic_url", "goods_number").
		Where("is_on_sale = ? and is_delete = ? and category_id = ?", 1, 0, Id).
		Order("sort_order asc").
		Scan(&goods)

	r1 := count / Size
	r2 := count % Size
	var div int

	if r2 > 0 {
		div = r1 + 1
	} else {
		div = r1
	}

	res = model.CategoryCurrentList{
		Count:       count,
		TotalPages:  div,
		PageSize:    Size,
		CurrentPage: Page,
		Data:        goods,
	}

	return
}
