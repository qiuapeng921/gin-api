package permissions

// Entity is the golang structure for table permissions.
type Entity struct {
	Id        int    `xorm:"not null pk autoincr comment('主键') INT(10)" json:"id"`
	Name      string `xorm:"not null default '' comment('权限名称') VARCHAR(30)" json:"name"`
	Url       string `xorm:"not null default '' comment('访问url地址') VARCHAR(50)" json:"url"`
	ParentId  int    `xorm:"default 0 comment('上级id') INT(10)" json:"parent_id"`
	Status    int    `xorm:"default 0 comment('是否可用 0：正常，1 禁用') TINYINT(1)" json:"status"`
	CreatedAt int    `xorm:"created not null default 0 comment('添加时间') INT(11)" json:"created_at"`
	UpdatedAt int    `xorm:"updated not null default 0 comment('修改时间') INT(11)" json:"updated_at"`
	DeletedAt int    `xorm:"deleted not null default 0 comment('删除时间') INT(11)" json:"deleted_at"`
}

func (Entity) TableName() string {
	return "permissions"
}
