package model

import "mall/internal/model/entity"

// index
type SearchIndex struct {
	DefaultKeyword     entity.HiolabsKeywords `json:"defaultKeyword"`
	HistoryKeywordList []string               `json:"historyKeywordList"`
	HotKeywordList     []HotKeywordList       `json:"hotKeywordList"`
}

type HotKeywordList struct {
	Keyword string `json:"keyword"`
	IsHot   int    `json:"is_hot"`
}
