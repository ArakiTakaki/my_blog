package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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

// SetPassword パスワードをセットする
func (u *User) SetPassword(password string) {
	hash, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if e != nil {
		panic(e)
	}
	u.Password = string(hash)
}

// CanLogin パスワードを比較するメソッド
func (u *User) CanLogin(password string) (uint, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return u.ID, err
}

// UserMeta ユーザの詳細情報
type UserMeta struct {
	gorm.Model
	UserID uint
	Key    string `gorm:"index"`
	Value  string
}
