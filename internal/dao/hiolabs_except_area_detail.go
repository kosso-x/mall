// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mall/internal/dao/internal"
)

// internalHiolabsExceptAreaDetailDao is internal type for wrapping internal DAO implements.
type internalHiolabsExceptAreaDetailDao = *internal.HiolabsExceptAreaDetailDao

// hiolabsExceptAreaDetailDao is the data access object for table hiolabs_except_area_detail.
// You can define custom methods on it to extend its functionality as you wish.
type hiolabsExceptAreaDetailDao struct {
	internalHiolabsExceptAreaDetailDao
}

var (
	// HiolabsExceptAreaDetail is globally public accessible object for table hiolabs_except_area_detail operations.
	HiolabsExceptAreaDetail = hiolabsExceptAreaDetailDao{
		internal.NewHiolabsExceptAreaDetailDao(),
	}
)

// Fill with you ideas below.
