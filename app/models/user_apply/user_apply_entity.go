package user_apply

type Entity struct {
	Id          int    `json:"id" xorm:"not null pk autoincr comment('用户申请添加好友表') INT(10)"`
	ApplyUserId int    `json:"apply_user_id" xorm:"not null default 0 comment('申请人ID') INT(10)"`
	UserId      int    `json:"user_id" xorm:"not null default 0 comment('添加人ID') INT(10)"`
	Message     string `json:"message" xorm:"not null default '' comment('申请消息') VARCHAR(255)"`
	Status      int    `json:"status" xorm:"not null default 0 comment('状态(0：待审核1：通过2：拒绝)') TINYINT(1)"`
	CreateTime  int    `json:"create_time" xorm:"not null default 0 comment('申请时间') INT(10)"`
	UpdateTime  int    `json:"update_time" xorm:"not null default 0 comment('修改时间') INT(10)"`
}

func (Entity) TableName() string {
	return "user_apply"
}