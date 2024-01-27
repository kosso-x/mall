// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// HiolabsAd is the golang structure for table hiolabs_ad.
type HiolabsAd struct {
	Id        uint   `json:"id"        ` //
	LinkType  int    `json:"linkType"  ` // 0商品，1链接
	Link      string `json:"link"      ` //
	GoodsId   int    `json:"goodsId"   ` //
	ImageUrl  string `json:"imageUrl"  ` //
	EndTime   int    `json:"endTime"   ` //
	Enabled   uint   `json:"enabled"   ` //
	SortOrder int    `json:"sortOrder" ` //
	IsDelete  int    `json:"isDelete"  ` //
}
