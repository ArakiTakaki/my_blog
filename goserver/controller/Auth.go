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
		c.JSON(http.StatusUnauthorized, gin.H{"status": "NG"})
		return
	}

	database := db.GetDB()
	defer database.Close()
	user := model.User{}
	database.Where("login = ?", json.User).First(&user)
	err = user.CanLogin(json.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "NG"})
		return
	}
	session := sessions.Default(c)
	session.Set(content.SessionAlive, true)
	session.Set(content.SessionUserID, user.ID)
	session.Save()
	c.JSON(http.StatusOK, gin.H{"status": "OK", "user_id": user.ID})
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
		c.JSON(http.StatusBadRequest, gin.H{"status": "NG"})
		return
	}
}
