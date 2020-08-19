package user

import "gin-api/app/models"

type Users struct {
	models.BaseModel
	Username   string `gorm:"type:VARCHAR(20);NOT NULL"`
	Password   string `gorm:"type:CHAR(60);NOT NULL"`
	Phone      string `gorm:"type:CHAR(11);NOT NULL"`
	Email      string `gorm:"type:VARCHAR(50);NOT NULL"`
	Avatar     string `gorm:"type:VARCHAR(255);NOT NULL"`
	Status     uint8  `gorm:"type:TINYINT(1) UNSIGNED;NOT NULL"`
}

