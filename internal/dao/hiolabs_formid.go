// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mall/internal/dao/internal"
)

// internalHiolabsFormidDao is internal type for wrapping internal DAO implements.
type internalHiolabsFormidDao = *internal.HiolabsFormidDao

// hiolabsFormidDao is the data access object for table hiolabs_formid.
// You can define custom methods on it to extend its functionality as you wish.
type hiolabsFormidDao struct {
	internalHiolabsFormidDao
}

var (
	// HiolabsFormid is globally public accessible object for table hiolabs_formid operations.
	HiolabsFormid = hiolabsFormidDao{
		internal.NewHiolabsFormidDao(),
	}
)

// Fill with you ideas below.
