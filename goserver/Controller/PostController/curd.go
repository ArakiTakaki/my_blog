package PostController

import (
	"net/http"
	"strconv"

	"github.com/ArakiTakaki/my_blog/goserver/db"
	"github.com/ArakiTakaki/my_blog/goserver/model"
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
	// data := db.GetDB()
	// defer db.Close()
	firstname := c.DefaultQuery("firstname", "Guest")
	c.JSON(http.StatusOK, firstname)
}

// Edit 記事の編集後のAPI
func Edit(c *gin.Context) {
	data := db.GetDB()
	defer data.Close()
	articleID := c.Param("articleID")
	posts := model.Post{}
	i, _ := strconv.Atoi(articleID)
	posts.ID = uint(i)
	// TODO ちゃんとアップデートするように変更する。
	data.Update(&posts)
	c.JSON(http.StatusOK, posts)

}
