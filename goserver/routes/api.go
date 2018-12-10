package routes

import (
	"net/http"

	"github.com/ArakiTakaki/my_blog/goserver/middleware"

	"github.com/ArakiTakaki/my_blog/goserver/controller"
	"github.com/gin-gonic/gin"
)

const ARTICLE_ID = "articleID"

// SetApi APIをセットする関数
func SetApi(r *gin.Engine) {
	r.POST("/sample", func(c *gin.Context) {
		c.JSON(http.StatusOK, "{hello:hello}")
	})
	r.LoadHTMLGlob("public/*")

	api := r.Group("api")

	// postsDB関連の操作
	posts := api.Group("posts")
	{
		posts.POST("/create", middleware.AuthUser, controller.PostCreate)
		posts.GET("/:articleID/edit", middleware.AuthUser, controller.PostEdit)
		posts.GET("/:articleID/show", middleware.AuthUser, controller.PostShow)
		posts.GET("/", controller.PostFind)
	}

	// auth関連の操作
	auth := api.Group("auth")
	{
		auth.POST("/login", controller.AuthLogin)
		auth.POST("/logout", middleware.AuthUser, controller.AuthLogout)
		auth.POST("/register", controller.AuthRegister)
	}

	// static files
	r.StaticFile("bundle.js", "./public/bundle.js")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
