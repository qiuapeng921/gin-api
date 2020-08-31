package admins

// Entity is the golang structure for table admins.
type Entity struct {
	Id        int    `xorm:"not null pk autoincr comment('管理员表') INT(10)" json:"id"`
	Username  string `xorm:"not null default '' comment('用户名') VARCHAR(30)" json:"username"`
	Password  string `xorm:"not null default '' comment('密码') VARCHAR(100)" json:"password"`
	Phone     string `xorm:"default '' comment('手机号') CHAR(11)" json:"phone"`
	Status    int    `xorm:"not null default 0 comment('状态:0正常 1禁用') TINYINT(1)" json:"status"`
	LoginIp   string `xorm:"default '' comment('最后登录ip') VARCHAR(15)" json:"login_ip"`
	LoginTime int    `xorm:"not null default 0 comment('最后登录时间') INT(11)" json:"login_time"`
	CreatedAt int    `xorm:"created not null default 0 comment('添加时间') INT(11)" json:"created_at"`
	UpdatedAt int    `xorm:"updated not null default 0 comment('修改时间') INT(11)" json:"updated_at"`
	DeletedAt int    `xorm:"deleted not null default 0 comment('删除时间') INT(11)" json:"deleted_at"`
}

func (Entity) TableName() string {
	return "admins"
}
