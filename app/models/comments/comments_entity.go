package comments

// Entity is the golang structure for table comments.
type Entity struct {
	Id        int    `xorm:"not null pk autoincr comment('ID') INT(10)" json:"id"`
	UserId    int    `xorm:"not null comment('发表用户ID') INT(10)" json:"user_id"`
	ArticleId int    `xorm:"not null comment('评论博文ID') INT(10)" json:"article_id"`
	Content   string `xorm:"not null comment('评论内容') TEXT" json:"content"`
	ParentId  int64  `xorm:"not null comment('父评论ID') BIGINT(20)" json:"parent_id"`
	CreatedAt int    `xorm:"created not null default 0 comment('添加时间') INT(11)" json:"created_at"`
	UpdatedAt int    `xorm:"updated not null default 0 comment('修改时间') INT(11)" json:"updated_at"`
	DeletedAt int    `xorm:"deleted not null default 0 comment('删除时间') INT(11)" json:"deleted_at"`
}

func (Entity) TableName() string {
	return "comments"
}
