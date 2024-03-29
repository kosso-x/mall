// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mall/internal/dao/internal"
)

// internalHiolabsFreightTemplateDao is internal type for wrapping internal DAO implements.
type internalHiolabsFreightTemplateDao = *internal.HiolabsFreightTemplateDao

// hiolabsFreightTemplateDao is the data access object for table hiolabs_freight_template.
// You can define custom methods on it to extend its functionality as you wish.
type hiolabsFreightTemplateDao struct {
	internalHiolabsFreightTemplateDao
}

var (
	// HiolabsFreightTemplate is globally public accessible object for table hiolabs_freight_template operations.
	HiolabsFreightTemplate = hiolabsFreightTemplateDao{
		internal.NewHiolabsFreightTemplateDao(),
	}
)

// Fill with you ideas below.
