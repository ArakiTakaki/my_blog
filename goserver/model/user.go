package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User ユーザーの基本的な情報
type User struct {
	gorm.Model
	Login       string `gorm:"index"`
	Password    string
	Age         int
	Birthday    time.Time
	DisplayName string `gorm:"size:255"`
	Email       string
	Name        string
	IgnoreMe    int        `gorm:"-"` // Ignore this field
	Meta        []UserMeta `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	Post        []Post     `gorm:"foreignkey:UserID;association_foreignkey:ID"`
}

// UserMeta ユーザの詳細情報
type UserMeta struct {
	gorm.Model
	UserID uint
	Key    string `gorm:"index"`
	Value  string
}
