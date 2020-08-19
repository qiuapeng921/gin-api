package user_behavior

// Entity is the golang structure for table user_behavior.
type Entity struct {
	Id              int    `xorm:"not null pk autoincr comment('主键id') INT(10)" json:"id"`
	UserId          int    `xorm:"not null default 0 comment('用户id') index INT(10)" json:"user_id"`
	DeviceType      string `xorm:"not null default '' comment('设备类型') VARCHAR(32)" json:"device_type"`
	PlatformType    string `xorm:"not null default '' comment('平台类型') VARCHAR(32)" json:"platform_type"`
	PlatformVersion string `xorm:"not null default '' comment('平台版本') VARCHAR(64)" json:"platform_version"`
	BrowserType     string `xorm:"not null default '' comment('浏览器类型') VARCHAR(32)" json:"browser_type"`
	BrowserVersion  string `xorm:"not null default '' comment('浏览器版本') VARCHAR(64)" json:"browser_version"`
	LoginIp         string `xorm:"not null default '' comment('登录ip') VARCHAR(50)" json:"login_ip"`
	CreatedAt       int    `xorm:"not null default 0 comment('添加时间') INT(11)" json:"created_at"`
	UpdatedAt       int    `xorm:"not null default 0 comment('修改时间') INT(11)" json:"updated_at"`
	DeletedAt       int    `xorm:"not null default 0 comment('删除时间') INT(11)" json:"deleted_at"`
}

func (Entity) TableName() string {
	return "user_behavior"
}
