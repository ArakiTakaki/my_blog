package middleware

import (
	"net/http"

	"github.com/ArakiTakaki/my_blog/goserver/content"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthUser 認証していない場合は弾く用途
// developmentだと無効化される
func AuthUser(c *gin.Context) {
	session := sessions.Default(c)
	if !content.SessionMode {
		session.Set(content.SessionAlive, true)
		session.Set(content.SessionUserID, uint(1))
	}
	alive := session.Get(content.SessionAlive)
	if alive == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "不正なページへのアクセスです"})
		c.Abort()
	}
	c.Next()
}
