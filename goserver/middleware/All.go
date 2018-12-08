package middleware

import (
	"net/http"
	"time"

	"github.com/ArakiTakaki/my_blog/goserver/content"
	"github.com/ArakiTakaki/my_blog/goserver/content/sessionContent"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SetMiddleware 全体に反映させる用途のmiddleware群
func SetMiddleware(r *gin.Engine) {
	// development middlewares
	if content.MODE == "development" {
		r.Use(Delay)
	}
	// production middlewares
	if content.MODE == "production" {
	}
}

// Delay 開発環境用、遅延をあえて出させるmiddleware
func Delay(c *gin.Context) {
	time.Sleep(1200 * time.Millisecond)
	c.Next()
}

// Auth 認証していない場合は弾く用途
func AuthUser(c *gin.Context) {
	alive := sessions.Default(c).Get(sessionContent.Alive)
	if alive == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "不正なページへのアクセスです"})
		c.Abort()
	}
	c.Next()
}
