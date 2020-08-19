package role_permission

// Entity is the golang structure for table role_permission.
type Entity struct {
    Id           int `xorm:"not null pk autoincr comment('主键id') INT(10)" json:"id"`
    RoleId       int `xorm:"not null comment('角色id') INT(10)" json:"role_id"`
    PermissionId int `xorm:"not null comment('权限id') INT(10)" json:"permission_id"`
}

func (Entity) TableName() string {
    return "role_permission"
}