package SessionController

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Register 会員登録する際のメソッド
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
