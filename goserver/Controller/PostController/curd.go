package PostController

import (
	"fmt"
	"net/http"
	"strconv"

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
	var err error
	post := model.Post{}
	// postDetail := model.PostDetail{}
	err = c.ShouldBindJSON(&post)
	// err = c.ShouldBindJSON(&postDetail)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSON compile error"})
		return
	}
	// session := sessions.Default(c)
	// post.UserID = session.Get(sessionContent.UserID).(uint)
	post.UserID = 0

	data := db.GetDB()
	defer data.Close()

	create := data.NewRecord(post)
	if !create {
		c.JSON(http.StatusBadRequest, gin.H{"message": "投稿できませんでした"})
		return
	}
	data.Create(&post)
	c.JSON(http.StatusOK, post)
}

// Edit 記事の編集後のAPI
func Edit(c *gin.Context) {
	data := db.GetDB()
	defer data.Close()
	session := sessions.Default(c)
	articleID, err := strconv.Atoi(c.Param("articleID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "不正なリクエストです"})
		return
	}
	post := model.Post{}
	userID := session.Get(sessionContent.UserID).(uint)
	data.First(&post, "id = ?", articleID)
	if post.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{"message": "不正なリクエストです"})
		return
	}
	err = c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "不正なリクエストです"})
		return
	}
	data.Update(&post)
	c.JSON(http.StatusOK, post)
}
