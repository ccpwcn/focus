// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// UserRole is the golang structure for table gf_user_role.
type UserRole struct {
    Id        uint        `orm:"id,primary" json:"id"`         // 自增ID    
    UserId    uint        `orm:"user_id"    json:"user_id"`    // 用户ID    
    RoleId    uint        `orm:"role_id"    json:"role_id"`    // 角色ID    
    CreatedAt *gtime.Time `orm:"created_at" json:"created_at"` // 创建时间  
}