package user_group_record

type Entity struct {
	Id         int64  `json:"id" xorm:"pk autoincr comment('用户组聊天记录表') BIGINT(20)"`
	GroupId    int    `json:"group_id" xorm:"not null comment('组ID') index(group_id) INT(10)"`
	UserId     int    `json:"user_id" xorm:"not null comment('用户ID') INT(10)"`
	Content    string `json:"content" xorm:"not null comment('消息内容') TINYTEXT"`
	Status     int    `json:"status" xorm:"not null default 1 comment('状态(1:正常)') index(group_id) TINYINT(1)"`
	CreateTime int    `json:"create_time" xorm:"not null default 0 comment('创建时间') index(group_id) INT(10)"`
}

func (Entity) TableName() string {
	return "user_group_record"
}