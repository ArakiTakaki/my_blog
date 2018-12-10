package middleware

import (
	"github.com/ArakiTakaki/my_blog/goserver/content"
	"github.com/gin-gonic/gin"
)

// SetMiddleware 全体に反映させる用途のmiddleware群
func SetMiddleware(r *gin.Engine) {
	// development middlewares
	if content.MODE == "development" {
		r.Use(Delay)
		r.Use(Cors)
	}
	// production middlewares
	if content.MODE == "production" {
	}
}

// Delay 開発環境用、遅延をあえて出させるmiddleware
func Delay(c *gin.Context) {
	// time.Sleep(1200 * time.Millisecond)

	c.Next()
}
