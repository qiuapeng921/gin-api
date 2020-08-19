package models

type BaseModel struct {
	Id         uint32 `gorm:"type:INT(10) UNSIGNED;AUTO_INCREMENT;NOT NULL"`
	CreateTime uint32 `gorm:"type:INT(11) UNSIGNED;NOT NULL"`
	UpdateTime uint32 `gorm:"type:INT(11) UNSIGNED;NOT NULL"`
}