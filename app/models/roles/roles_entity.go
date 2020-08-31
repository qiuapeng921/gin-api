package roles

// Entity is the golang structure for table roles.
type Entity struct {
	Id        int    `xorm:"not null pk autoincr comment('角色表') INT(10)" json:"id"`
	RoleName  string `xorm:"not null default '' comment('角色名称') VARCHAR(30)" json:"role_name"`
	RoleDesc  string `xorm:"not null default '' comment('角色描述') VARCHAR(100)" json:"role_desc"`
	Status    int    `xorm:"not null default 0 comment('状态0 可用 1 禁用') TINYINT(1)" json:"status"`
	CreatedAt int    `xorm:"created not null default 0 comment('添加时间') INT(11)" json:"created_at"`
	UpdatedAt int    `xorm:"updated not null default 0 comment('修改时间') INT(11)" json:"updated_at"`
	DeletedAt int    `xorm:"deleted not null default 0 comment('删除时间') INT(11)" json:"deleted_at"`
}

func (Entity) TableName() string {
	return "roles"
}
