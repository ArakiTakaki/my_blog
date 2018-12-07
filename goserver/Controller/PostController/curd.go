package PostController

import (
	"github.com/ArakiTakaki/my_blog/goserver/db"
	"github.com/ArakiTakaki/my_blog/goserver/model"
	"github.com/gin-gonic/gin"
)

// LIMIT 記事の表示限界
const GET_LIMIT = 10

// 記事一覧を取得する
func All(c *gin.Context) {
	data := db.GetDB()
	defer data.Close()
	posts := []model.Post{}
	data.Find(&posts).Limit(GET_LIMIT)
	c.JSON(200, posts)
}
