package routes

import (
	"net/http"

	"github.com/ArakiTakaki/my_blog/goserver/middleware"

	"github.com/ArakiTakaki/my_blog/goserver/Controller/PostController"
	"github.com/ArakiTakaki/my_blog/goserver/Controller/SessionController"
	"github.com/gin-gonic/gin"
)

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
		posts.POST("/create", PostController.Create)
		// posts.GET("/:articleID/edit", middleware.AuthUser, PostController.Edit)
		posts.GET("/", PostController.All)
	}

	// auth関連の操作
	auth := api.Group("auth")
	{
		auth.POST("/login", SessionController.LoginController)
		auth.POST("/logout", middleware.AuthUser, SessionController.Logout)
		auth.POST("/register", SessionController.Register)
	}

	// static files
	r.StaticFile("bundle.js", "./public/bundle.js")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
