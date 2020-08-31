package user_friend

type Entity struct {
	Id       int `json:"id" xorm:"not null pk autoincr comment('好友表') INT(10)"`
	FriendId int `json:"friend_id" xorm:"not null comment('好友的ID') index(friend_id) INT(10)"`
	UserId   int `json:"user_id" xorm:"not null comment('我的ID') index(friend_id) INT(10)"`
	Status   int `json:"status" xorm:"not null default 0 comment('状态 0 正常 1 删除') TINYINT(1)"`
}

func (Entity) TableName() string {
	return "user_friend"
}