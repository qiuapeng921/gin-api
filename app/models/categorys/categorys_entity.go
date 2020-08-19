package categorys

// Entity is the golang structure for table categorys.
type Entity struct {
	Id          int    `xorm:"not null pk autoincr INT(10)" json:"id"`
	Name        string `xorm:"not null default '' comment('分类名称') VARCHAR(20)" json:"name"`
	Description string `xorm:"comment('分类描述') VARCHAR(50)" json:"description"`
	Status      int    `xorm:"not null default 0 comment('0 正常 1 删除') TINYINT(1)" json:"status"`
	CreatedAt   int    `xorm:"not null default 0 comment('添加时间') INT(11)" json:"created_at"`
	UpdatedAt   int    `xorm:"not null default 0 comment('修改时间') INT(11)" json:"updated_at"`
	DeletedAt   int    `xorm:"not null default 0 comment('删除时间') INT(11)" json:"deleted_at"`
}

func (Entity) TableName() string {
	return "categorys"
}
