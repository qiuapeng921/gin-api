package user_record

type Entity struct {
	Id         int    `json:"id" xorm:"not null pk autoincr comment('用户聊天记录表') INT(10)"`
	UserId     int    `json:"user_id" xorm:"not null comment('用户ID') INT(10)"`
	Content    string `json:"content" xorm:"not null comment('消息内容') TINYTEXT"`
	Status     int    `json:"status" xorm:"not null default 1 comment('状态(1：正常)') TINYINT(1)"`
	CreateTime int    `json:"create_time" xorm:"not null default 0 comment('创建时间') INT(10)"`
}

func (Entity) TableName() string {
	return "user_record"
}
