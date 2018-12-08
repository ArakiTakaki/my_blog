package model

import (
	"github.com/jinzhu/gorm"
)

// Post 記事の投稿の基盤となるクラス
type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Excerpt string `json:"excerpt"`
	// open: 公開状態, close: 閲覧禁止, temp: 一時保存, limited : 限定公開
	Status        string     `json:"status" gorm:"index"`
	CommentStatus string     `json:"comment_status"`
	Password      string     `json:"password"`
	UserID        uint       `json:"-"`
	Detail        PostDetail `json:"defail"`
}

// PostDetail 投稿された記事の詳細
type PostDetail struct {
	gorm.Model
	Content string    `json:"content"`
	Comment []Comment `json:"comments"`
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
