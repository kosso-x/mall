// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mall/internal/dao/internal"
)

// internalHiolabsKeywordsDao is internal type for wrapping internal DAO implements.
type internalHiolabsKeywordsDao = *internal.HiolabsKeywordsDao

// hiolabsKeywordsDao is the data access object for table hiolabs_keywords.
// You can define custom methods on it to extend its functionality as you wish.
type hiolabsKeywordsDao struct {
	internalHiolabsKeywordsDao
}

var (
	// HiolabsKeywords is globally public accessible object for table hiolabs_keywords operations.
	HiolabsKeywords = hiolabsKeywordsDao{
		internal.NewHiolabsKeywordsDao(),
	}
)

// Fill with you ideas below.
