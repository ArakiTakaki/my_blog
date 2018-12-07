package SessionController

import (
	"fmt"
	"net/http"

	"github.com/ArakiTakaki/my_blog/goserver/db"

	"github.com/ArakiTakaki/my_blog/goserver/model"
	"github.com/gin-gonic/gin"
)

// Register 会員登録する際のメソッド
func Register(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "NG"})
		return
	}

	user.HashToPassword()
	database := db.GetDB()
	defer database.Close()
	var count int
	database.Model(&model.User{}).Where("login = ?", user.Login).Count(&count)
	fmt.Println(count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ユーザが作成できませんでした"})
		return
	}
	database.Create(&user)
	c.JSON(http.StatusBadRequest, gin.H{"message": "ユーザの作成が完了しました"})

}
