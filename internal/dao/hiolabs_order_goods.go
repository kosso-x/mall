// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mall/internal/dao/internal"
)

// internalHiolabsOrderGoodsDao is internal type for wrapping internal DAO implements.
type internalHiolabsOrderGoodsDao = *internal.HiolabsOrderGoodsDao

// hiolabsOrderGoodsDao is the data access object for table hiolabs_order_goods.
// You can define custom methods on it to extend its functionality as you wish.
type hiolabsOrderGoodsDao struct {
	internalHiolabsOrderGoodsDao
}

var (
	// HiolabsOrderGoods is globally public accessible object for table hiolabs_order_goods operations.
	HiolabsOrderGoods = hiolabsOrderGoodsDao{
		internal.NewHiolabsOrderGoodsDao(),
	}
)

// Fill with you ideas below.