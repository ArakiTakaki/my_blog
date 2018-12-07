package routes

import (
	"net/http"

	"github.com/ArakiTakaki/my_blog/goserver/Controller/PostController"
	"github.com/gin-gonic/gin"
)

// SetApi APIをセットする関数
func SetApi(r *gin.Engine) {
	r.POST("/sample", func(c *gin.Context) {
		c.JSON(http.StatusOK, "{hello:hello}")
	})
	r.LoadHTMLGlob("public/*")

	api := r.Group("api")
	posts := api.Group("posts")
	{
		posts.GET("/", PostController.All)

	}
	// static files
	r.StaticFile("bundle.js", "./public/bundle.js")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
