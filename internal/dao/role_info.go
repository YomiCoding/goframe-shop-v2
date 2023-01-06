// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goframe-shop-v2/internal/dao/internal"
)

// internalRoleInfoDao is internal type for wrapping internal DAO implements.
type internalRoleInfoDao = *internal.RoleInfoDao

// roleInfoDao is the data access object for table role_info.
// You can define custom methods on it to extend its functionality as you wish.
type roleInfoDao struct {
	internalRoleInfoDao
}

var (
	// RoleInfo is globally public accessible object for table role_info operations.
	RoleInfo = roleInfoDao{
		internal.NewRoleInfoDao(),
	}
)

// Fill with you ideas below.
