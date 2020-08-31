package admin_role


// Entity is the golang structure for table admin_role.
type Entity struct {
    Id      int `xorm:"not null pk autoincr comment('主键i用户角色关系表') INT(10)" json:"id"`
    AdminId int `xorm:"not null comment('管理员id') INT(10)" json:"admin_id"`
    RoleId  int `xorm:"not null comment('角色id') INT(10)" json:"role_id"`
}

func (Entity) TableName() string {
    return "admin_role"
}