package article_tags

// Entity is the golang structure for table article_tags.
type Entity struct {
	Id        int `xorm:"not null pk INT(10)" json:"id"`
	ArticleId int `xorm:"INT(10)" json:"article_id"`
	TagId     int `xorm:"INT(10)" json:"tag_id"`
}

func (Entity) TableName() string {
	return "article_tags"
}
