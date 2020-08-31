package user_group_member

type Entity struct {
	Id      int `json:"id" xorm:"not null pk autoincr comment('用户组成员表') INT(10)"`
	UserId  int `json:"user_id" xorm:"not null comment('用户ID') index(member_id) INT(10)"`
	GroupId int `json:"group_id" xorm:"not null comment('组ID') index(member_id) INT(10)"`
}

func (Entity) TableName() string {
	return "user_group_member"
}