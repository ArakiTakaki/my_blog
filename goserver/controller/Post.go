package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ArakiTakaki/my_blog/goserver/content"

	"github.com/ArakiTakaki/my_blog/goserver/db"
	"github.com/ArakiTakaki/my_blog/goserver/model"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// limit 1度に表示する記事の数
const limit = 10

// Viewさせる文字数
const excerpt = 50

const firstPage = 1

var lastPage = 0

// PostFind 記事一覧を取得する
func PostFind(c *gin.Context) {
	query := c.Query("p")
	offset, location := pagination(query)
	data := db.GetDB()
	defer data.Close()
	posts := []model.Post{}
	data.Limit(limit).Offset(offset).Order("created_at DESC").Find(&posts, "status = ?", "open")
	if lastPage == 0 {
		data.Model(&model.Post{}).Count(&lastPage)
		fmt.Println(lastPage)
		if lastPage%limit == 0 {
			lastPage = lastPage / limit
		} else {
			lastPage = lastPage/limit + 1
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"post_content": posts,
		"location":     location,
		"firstPage":    firstPage,
		"lastPage":     lastPage,
	})
}

func pagination(query string) (int, int) {
	var err error
	location, err := strconv.Atoi(query)
	if err != nil {
		return 0, 1
	}
	offset := ((location - 1) * limit)
	return offset, location
}

// PostCreate 記事の作成
func PostCreate(c *gin.Context) {
	var err error
	post := model.Post{}
	err = c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSON compile error"})
		return
	}
	session := sessions.Default(c)
	post.UserID = session.Get(content.SessionUserID).(uint)

	data := db.GetDB()
	defer data.Close()

	create := data.NewRecord(post)
	if !create {
		c.JSON(http.StatusBadRequest, gin.H{"message": "投稿できませんでした"})
		return
	}
	post.Excerpt = CreateExcerpt(post.Detail.Content)
	data.Create(&post)
	c.JSON(http.StatusOK, post)
}

type editContent struct {
	Content string `json:"content" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Status  string `json:"status" binding:"required"`
	Detail  model.PostDetail
}

// PostEdit 記事の編集後のAPI
func PostEdit(c *gin.Context) {
	var err error
	var post model.Post
	data := db.GetDB()
	defer data.Close()
	session := sessions.Default(c)
	post.UserID = session.Get(content.SessionUserID).(uint)
	data.First(&post, c.Param("articleID"))
	var ec editContent
	err = c.ShouldBindJSON(&ec)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSON compile error"})
		return
	}
	post.Excerpt = CreateExcerpt(ec.Detail.Content)
	post.Detail = ec.Detail
	post.Status = ec.Status
	post.Title = ec.Title
	post.UpdatedAt = time.Now()
	data.Update(&post)
	c.JSON(http.StatusOK, post)
}

// PostShow 内容を投稿する
func PostShow(c *gin.Context) {
	data := db.GetDB()
	defer data.Close()
	articleID := c.Param("articleID")
	if articleID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	pd := model.PostDetail{}
	data.First(&pd, articleID)
	c.JSON(http.StatusOK, pd)
}

// CreateExcerpt 短縮文字列を作成する。
func CreateExcerpt(data string) string {
	text := []byte(data)
	return string(text[:excerpt]) + "..."
}
