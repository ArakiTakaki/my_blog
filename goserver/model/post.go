package model

import (
	"github.com/jinzhu/gorm"
)

// Post 記事の投稿の基盤となるクラス
type Post struct {
	gorm.Model
	Title         string
	Excerpt       string
	Status        string
	CommentStatus string
	Password      string
	UserID        uint
	Detail        PostDetail
}

// PostDetail 投稿された記事の詳細
type PostDetail struct {
	gorm.Model
	Content string
	Comment []Comment
}

// Comment 記事に投下されたコメント
type Comment struct {
	gorm.Model
	DisplayName string
	URL         string
	Comment     string
	Meta        []CommentMeta `gorm:"foreignkey:CommentID;association_foreignkey:ID"`
}

// CommentMeta コメントのメタ要素
type CommentMeta struct {
	gorm.Model
	CommentID uint
	Key       string `gorm:"index"`
	Value     string
}
