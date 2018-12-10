package controller

import (
	"net/http"

	"github.com/ArakiTakaki/my_blog/goserver/content"
	"github.com/ArakiTakaki/my_blog/goserver/db"
	"github.com/ArakiTakaki/my_blog/goserver/model"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Binding from JSON
type loginModel struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// AuthLogin ログインを実装
func AuthLogin(c *gin.Context) {
	var json loginModel
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "NG"})
		return
	}

	database := db.GetDB()
	defer database.Close()
	user := model.User{}
	database.Where("login = ?", json.User).First(&user)
	err = user.CanLogin(json.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "NG"})
		return
	}
	session := sessions.Default(c)
	session.Set(content.SessionAlive, true)
	session.Set(content.SessionUserID, user.ID)
	session.Save()
	c.JSON(http.StatusOK, gin.H{"messgage": "OK", "user_id": user.ID, "user_name": user.DisplayName, "full_name": user.Name})
}

// AuthLogout 会員登録する際のメソッド
func AuthLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

// AuthRegister 会員登録する際のメソッド
func AuthRegister(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "NG"})
		return
	}
	user.PasswordToHash()

	data := db.GetDB()
	defer data.Close()
	result := data.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"message": "NG"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// AuthEdit 会員情報を変更する際のメソッド
func AuthEdit(c *gin.Context) {
	// TODO 必要になったら記載する
	return
}
