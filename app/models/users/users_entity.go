package users

type Entity struct {
	Id        int    `xorm:"not null pk autoincr INT(10)" json:"id"`
	Username  string `xorm:"not null default '' comment('用户名') unique VARCHAR(20)" json:"username"`
	Password  string `xorm:"not null default '' comment('密码') CHAR(60)" json:"password"`
	Phone     string `xorm:"not null default '' comment('手机号') unique CHAR(11)" json:"phone"`
	Email     string `xorm:"not null default '' comment('邮箱') index VARCHAR(50)" json:"email"`
	Avatar    string `xorm:"not null default '' comment('头像') VARCHAR(255)" json:"avatar"`
	Status    int    `xorm:"not null default 1 comment('状态 0:删除 1:正常 2:禁用') TINYINT(1)" json:"status"`
	CreatedAt int    `xorm:"created not null default 0 comment('添加时间') INT(11)" json:"created_at"`
	UpdatedAt int    `xorm:"updated not null default 0 comment('修改时间') INT(11)" json:"updated_at"`
	DeletedAt int    `xorm:"deleted not null default 0 comment('删除时间') INT(11)" json:"deleted_at"`
}

func (Entity) TableName() string {
	return "users"
}
