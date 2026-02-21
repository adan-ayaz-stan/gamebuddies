package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model

	Title    string `gorm:"type:varchar(255);not null"`
	Subtitle string `gorm:"type:varchar(255)"`
	ImgURL   string `gorm:"type:varchar(255);not null;column:img_url"`
}
