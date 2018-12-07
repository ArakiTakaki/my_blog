package SessionController

import (
	"net/http"

	"github.com/ArakiTakaki/my_blog/goserver/model"
	"github.com/gin-gonic/gin"
)

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func LoginController(c *gin.Context) {
	var json Login
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "NG"})
		return
	}

	// database := db.GetDB()
	// defer database.Close()
	user := model.User{}
	user.Login = "hogehoge"
	user.Password = "123"

	// database.Where("login = ?", json.User).First(&user)
	if user.Password != json.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "NG"})
		return
	}

	// session := sessions.Default(c)
	// session.Set("alive", true)
	// session.Set("userID", user.ID)
	// session.Save()
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
