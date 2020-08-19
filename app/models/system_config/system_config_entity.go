package system_config

// Entity is the golang structure for table system_config.
type Entity struct {
	Id        int    `xorm:"not null pk autoincr INT(10)" json:"id"`
	Key       string `xorm:"not null comment('键') VARCHAR(50)" json:"key"`
	Value     int    `xorm:"not null comment('值') TINYINT(255)" json:"value"`
	Desc      string `xorm:"comment('描述') VARCHAR(255)" json:"desc"`
	CreatedAt int    `xorm:"not null default 0 comment('添加时间') INT(11)" json:"created_at"`
	UpdatedAt int    `xorm:"not null default 0 comment('修改时间') INT(11)" json:"updated_at"`
	DeletedAt int    `xorm:"not null default 0 comment('删除时间') INT(11)" json:"deleted_at"`
}

func (Entity) TableName() string {
	return "system_config"
}
