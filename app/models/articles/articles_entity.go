package articles

// Entity is the golang structure for table articles.
type Entity struct {
	Id           int    `xorm:"not null pk autoincr comment('主键id') INT(10)" json:"id"`
	UserId       int    `xorm:"not null comment('用户ID') INT(10)" json:"user_id"`
	Title        string `xorm:"not null comment('文章标题') TEXT" json:"title"`
	CategoryId   int    `xorm:"not null default 0 comment('分类id') INT(10)" json:"category_id"`
	Content      string `xorm:"not null comment('博文内容') LONGTEXT" json:"content"`
	Views        int64  `xorm:"not null default 0 comment('浏览量') BIGINT(20)" json:"views"`
	LikeCount    int64  `xorm:"not null default 0 comment('喜欢数') BIGINT(20)" json:"like_count"`
	CommentCount int64  `xorm:"not null default 0 comment('评论总数') BIGINT(20)" json:"comment_count"`
	Status       int    `xorm:"not null default 0 comment('0 正常 1 删除') TINYINT(1)" json:"status"`
	CreatedAt    int    `xorm:"created not null default 0 comment('添加时间') INT(11)" json:"created_at"`
	UpdatedAt    int    `xorm:"updated not null default 0 comment('修改时间') INT(11)" json:"updated_at"`
	DeletedAt    int    `xorm:"deleted not null default 0 comment('删除时间') INT(11)" json:"deleted_at"`
}

func (Entity) TableName() string {
	return "articles"
}
