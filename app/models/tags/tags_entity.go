package tags

// Entity is the golang structure for table tags.
type Entity struct {
	Id          int    `xorm:"not null pk autoincr INT(11)" json:"id"`
	Name        string `xorm:"not null comment('标签名称') VARCHAR(50)" json:"name"`
	Description string `xorm:"not null comment('标签描述') VARCHAR(100)" json:"description"`
	Status      int    `xorm:"not null default 0 comment('0 正常 1 删除') TINYINT(1)" json:"status"`
	CreatedAt   int    `xorm:"created not null default 0 comment('添加时间') INT(11)" json:"created_at"`
	UpdatedAt   int    `xorm:"updated not null default 0 comment('修改时间') INT(11)" json:"updated_at"`
	DeletedAt   int    `xorm:"deleted not null default 0 comment('删除时间') INT(11)" json:"deleted_at"`
}

func (Entity) TableName() string {
	return "tags"
}
