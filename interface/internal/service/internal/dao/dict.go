// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"interface/internal/service/internal/dao/internal"
)

// dictDao is the data access object for table s_dict.
// You can define custom methods on it to extend its functionality as you wish.
type dictDao struct {
	*internal.DictDao
}

var (
	// Dict is globally public accessible object for table s_dict operations.
	Dict = dictDao{
		internal.NewDictDao(),
	}
)

// Fill with you ideas below.
