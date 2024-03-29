// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mall/internal/dao/internal"
)

// internalHiolabsSpecificationDao is internal type for wrapping internal DAO implements.
type internalHiolabsSpecificationDao = *internal.HiolabsSpecificationDao

// hiolabsSpecificationDao is the data access object for table hiolabs_specification.
// You can define custom methods on it to extend its functionality as you wish.
type hiolabsSpecificationDao struct {
	internalHiolabsSpecificationDao
}

var (
	// HiolabsSpecification is globally public accessible object for table hiolabs_specification operations.
	HiolabsSpecification = hiolabsSpecificationDao{
		internal.NewHiolabsSpecificationDao(),
	}
)

// Fill with you ideas below.
