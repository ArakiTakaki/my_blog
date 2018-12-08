package PostController

import (
	"errors"
	"net/http"

	"github.com/ArakiTakaki/my_blog/goserver/content/sessionContent"
	"github.com/ArakiTakaki/my_blog/goserver/db"
	"github.com/ArakiTakaki/my_blog/goserver/model"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// getLimit 記事の表示限界
const getLimit = 10

// All 記事一覧を取得する
func All(c *gin.Context) {
	data := db.GetDB()
	defer data.Close()
	posts := []model.Post{}
	data.Where("status = ?", "open").Limit(getLimit).Order("created_at DESC").Find(&posts)
	c.JSON(http.StatusOK, posts)
}

// Create 記事の作成
func Create(c *gin.Context) {
	post := model.Post{}
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "投稿できませんでした"})
		return
	}

	session := sessions.Default(c)
	post.UserID = session.Get(sessionContent.UserID).(uint)
	// data := db.GetDB()
	// defer db.Close()
	firstname := c.DefaultQuery("firstname", "Guest")
	c.JSON(http.StatusOK, firstname)
}

// Edit 記事の編集後のAPI
func Edit(c *gin.Context) {
	var err error
	data := db.GetDB()
	defer data.Close()
	session := sessions.Default(c)
	articleID := c.Param("articleID")
	post := model.Post{}
	userID := session.Get(sessionContent.UserID).(uint)
	data.Where("id = ?", articleID).First(&post)

	// validations
	err = c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "記事の形式が不正です"})
		return
	}
	err = canPostChange(post, userID)
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"message": "error"})
		return
	}
	data.Update(&post)
	c.JSON(http.StatusOK, post)
}

func canPostChange(post model.Post, userID uint) error {
	if post.UserID != userID {
		return errors.New("認証が失敗しました")
	}
	return nil
}
