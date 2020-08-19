package menus

// Entity is the golang structure for table menus.
type Entity struct {
	Id        int    `xorm:"not null pk autoincr comment('主键id') INT(10)" json:"id"`
	Name      string `xorm:"not null default '' comment('菜单名') index VARCHAR(20)" json:"name"`
	Url       string `xorm:"not null default '' comment('地址') VARCHAR(50)" json:"url"`
	Identify  string `xorm:"not null default '' comment('标识') VARCHAR(20)" json:"identify"`
	ParentId  int    `xorm:"not null default 0 comment('上级id') INT(10)" json:"parent_id"`
	Icon      string `xorm:"not null default '' comment('图标') VARCHAR(50)" json:"icon"`
	Status    int    `xorm:"not null default 1 comment('状态 1 正常 0 删除') TINYINT(1)" json:"status"`
	CreatedAt int    `xorm:"not null default 0 comment('添加时间') INT(11)" json:"created_at"`
	UpdatedAt int    `xorm:"not null default 0 comment('修改时间') INT(11)" json:"updated_at"`
	DeletedAt int    `xorm:"not null default 0 comment('删除时间') INT(11)" json:"deleted_at"`
}

func (Entity) TableName() string {
	return "menus"
}
