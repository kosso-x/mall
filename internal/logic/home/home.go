package home

import (
	"context"
	"mall/internal/dao"
	"mall/internal/model"
	"mall/internal/service"

	"github.com/gogf/gf/errors/gerror"
)

type (
	sHome struct{}
)

func init() {
	service.RegisterHome(New())
}

func New() service.IHome {
	return &sHome{}
}

func (s *sHome) Info(ctx context.Context) (res *model.Info, err error) {
	r1, _ := s.InfoShortCategory(ctx)
	// fmt.Printf("r1: %v\n", r1)

	r2, _ := s.InfoShortAd(ctx)
	// fmt.Printf("r2: %v\n", r2)

	r3, _ := s.InfoShortNotice(ctx)
	// fmt.Printf("r3: %v\n", r3)

	r4, _ := s.InfoCategoryGood(ctx)
	// fmt.Printf("r4: %v\n", r4)

	r5, _ := s.InfoCartCount(ctx)
	// fmt.Printf("r5: %v\n", r5)

	res = &model.Info{
		Channel:      r1,
		Banner:       r2,
		Notice:       r3,
		CategoryList: r4,
		CartCount:    r5,
	}

	return
}

func (s *sHome) InfoShortCategory(ctx context.Context) (res []model.ShortCategory, err error) {
	err = dao.
		HiolabsCategory.
		Ctx(ctx).
		Fields("id", "icon_url", "name", "sort_order").
		Where("is_channel = ? and parent_id = ?", 1, 0).
		Order("sort_order asc").
		Scan(&res)

	if err != nil {
		return nil, gerror.New("get short-categories error")
	}

	return
}

func (s *sHome) InfoShortAd(ctx context.Context) (res []model.ShortAd, err error) {
	err = dao.
		HiolabsAd.
		Ctx(ctx).
		Fields("link_type", "goods_id", "image_url", "link").
		Where("enabled = ? and is_delete = ?", 1, 0).
		Order("sort_order asc").
		Scan(&res)

	if err != nil {
		return nil, gerror.New("get short-ads error")
	}

	return
}

func (s *sHome) InfoShortNotice(ctx context.Context) (res []model.ShortNotice, err error) {
	err = dao.
		HiolabsNotice.
		Ctx(ctx).
		Fields("content").
		Where("is_delete = ?", 0).
		Order("sort_order asc").
		Scan(&res)

	if err != nil {
		return nil, gerror.New("get short-notices error")
	}

	return
}

func (s *sHome) InfoCategoryGood(ctx context.Context) (res []model.CategoryGood, err error) {
	err = dao.
		HiolabsCategory.
		Ctx(ctx).
		Fields("id", "name", "img_url", "p_height").
		Where("parent_id = ? and is_show = ?", 0, 1).
		Order("sort_order asc").
		Scan(&res)

	if err != nil {
		return nil, gerror.New("get short-notices error")
	}

	for i := 0; i < len(res); i++ {
		goods, _ := s.InfoGoods(ctx, res[i].Id)
		res[i].GoodsList = goods
	}

	return
}

func (s *sHome) InfoGoods(ctx context.Context, category_id int) (res []model.GoodsList, err error) {
	err = dao.
		HiolabsGoods.
		Ctx(ctx).
		Fields("id", "list_pic_url", "is_new", "goods_number", "name", "min_retail_price").
		Where("category_id = ? and goods_number >= ? and is_on_sale = ? and is_index = ? and is_delete = ?", category_id, 0, 1, 1, 0).
		Order("sort_order asc").
		Scan(&res)

	if err != nil {
		return nil, gerror.New("get category-goods error")
	}

	return
}

func (s *sHome) InfoCartCount(ctx context.Context) (res int, err error) {
	err = dao.
		HiolabsCart.
		Ctx(ctx).
		Fields("SUM(number)").
		Where("user_id = ? and is_delete = ?", 0, 0).
		Scan(&res)

	if err != nil {
		return 0, gerror.New("get carts-count error")
	}

	return
}
