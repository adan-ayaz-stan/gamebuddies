package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model

	UserID uint `gorm:"unique;not null"`

	User *User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	DisplayName  string
	AvatarURL    string `gorm:"type:varchar(512)"`
	TotalMatches int    `gorm:"default:0"`
}
