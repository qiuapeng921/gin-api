package user_group

type Entity struct {
	Id         int    `json:"id" xorm:"not null pk autoincr comment('用户组表') INT(10)"`
	UserId     int    `json:"user_id" xorm:"not null comment('创建人用户ID') INT(10)"`
	GroupName  string `json:"group_name" xorm:"not null comment('组名') VARCHAR(20)"`
	Status     int    `json:"status" xorm:"not null default 1 comment('状态(1:正常)') TINYINT(1)"`
	CreateTime int    `json:"create_time" xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime int    `json:"update_time" xorm:"not null default 0 comment('更新时间') INT(10)"`
}

func (Entity) TableName() string {
	return "user_group"
}