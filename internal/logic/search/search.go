package search

import (
	"context"
	"mall/internal/cmd"
	"mall/internal/dao"
	"mall/internal/model"
	"mall/internal/model/entity"
	"mall/internal/service"

	"github.com/gogf/gf/errors/gerror"
)

type (
	sSearch struct{}
)

func init() {
	service.RegisterSearch(New())
}

func New() service.ISearch {
	return &sSearch{}
}

func (s *sSearch) Index(ctx context.Context) (res *model.SearchIndex, err error) {
	var (
		defaultKeyword     entity.HiolabsKeywords
		history            []entity.HiolabsKeywords
		historyKeywordList []string
		hotKeywordList     []model.HotKeywordList
	)

	err = dao.HiolabsKeywords.Ctx(ctx).Where("is_default = ?", 1).Scan(&defaultKeyword)
	if err != nil {
		return
	}
	err = dao.HiolabsKeywords.Ctx(ctx).Fields("DISTINCT (keyword)", "is_hot").Limit(0, 10).Scan(&hotKeywordList)
	if err != nil {
		return
	}
	err = dao.HiolabsSearchHistory.Ctx(ctx).Fields("DISTINCT keyword").Where("user_id = ?", cmd.CurrentUser.Id).Limit(0, 10).Scan(&history)
	if err != nil {
		return
	}
	for _, h := range history {
		historyKeywordList = append(historyKeywordList, h.Keyword)
	}
	res = &model.SearchIndex{
		DefaultKeyword:     defaultKeyword,
		HistoryKeywordList: historyKeywordList,
		HotKeywordList:     hotKeywordList,
	}

	return
}

func (s *sSearch) Helper(ctx context.Context, keyword string) (helpKeywords []string, err error) {
	var h_keywords []entity.HiolabsKeywords
	err = dao.HiolabsKeywords.Ctx(ctx).Where("keyword like ?", "%"+keyword+"%").Scan(&h_keywords)

	if err != nil {
		return
	}
	for _, h := range h_keywords {
		helpKeywords = append(helpKeywords, h.Keyword)
	}

	return
}

func (s *sSearch) ClearHistory(ctx context.Context) (err error) {
	_, err = dao.HiolabsSearchHistory.Ctx(ctx).Delete("user_id = ?", cmd.CurrentUser.Id)
	if err != nil {
		return gerror.New("delete hiolabs_search_history error")
	}

	return
}
