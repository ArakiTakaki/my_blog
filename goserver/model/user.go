package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User ユーザーの基本的な情報
type User struct {
	gorm.Model
	Login       string     `json:"login" gorm:"unique_index;not null" binding:"required"`
	Password    string     `json:"password" binding:"required"`
	Age         int        `json:"age"`
	Birthday    time.Time  `json:"birthday"`
	DisplayName string     `json:"display_name" gorm:"size:255"`
	Email       string     `json:"email"`
	Name        string     `json:"name"`
	IgnoreMe    int        `json:"-" gorm:"-"` // Ignore this field
	Meta        []UserMeta `json:"meta" gorm:"foreignkey:UserID;association_foreignkey:ID"`
	Posts       []Post     `json:"posts" gorm:"foreignkey:UserID;association_foreignkey:ID"`
}

// HashToPassword パスワードをセットする
func (u *User) HashToPassword() {
	hash, e := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if e != nil {
		panic(e)
	}
	u.Password = string(hash)
}

// CanLogin パスワードを比較するメソッド
func (u *User) CanLogin(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err
}

// UserMeta ユーザの詳細情報
type UserMeta struct {
	gorm.Model
	UserID uint
	Key    string `gorm:"index"`
	Value  string
}
