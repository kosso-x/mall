// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mall/internal/dao/internal"
)

// internalHiolabsGoodsSpecificationDao is internal type for wrapping internal DAO implements.
type internalHiolabsGoodsSpecificationDao = *internal.HiolabsGoodsSpecificationDao

// hiolabsGoodsSpecificationDao is the data access object for table hiolabs_goods_specification.
// You can define custom methods on it to extend its functionality as you wish.
type hiolabsGoodsSpecificationDao struct {
	internalHiolabsGoodsSpecificationDao
}

var (
	// HiolabsGoodsSpecification is globally public accessible object for table hiolabs_goods_specification operations.
	HiolabsGoodsSpecification = hiolabsGoodsSpecificationDao{
		internal.NewHiolabsGoodsSpecificationDao(),
	}
)

// Fill with you ideas below.
