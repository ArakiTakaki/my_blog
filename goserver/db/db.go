package db

import (
	"time"

	"github.com/ArakiTakaki/my_blog/goserver/model"
	"github.com/jinzhu/gorm"
)

// GetDB BDオブジェクトを取得する
func GetDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	user := model.User{}
	db.First(&user)
	return db
}

// Migration データベースの必要なテーブルを定義
func Migration(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UserMeta{})
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.PostDetail{})
	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.CommentMeta{})
}

// User 初期設定
func User(db *gorm.DB) {
	user := model.User{}
	user.ID = 1
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Age = 20
	user.DisplayName = "araki"
	user.Email = "sjyyj008@gmail.com"
	user.Birthday = time.Now()
	user.Name = "araki takkai"
	user.Password = "aiefjaskljfdlk"
	db.Create(&user)
}
