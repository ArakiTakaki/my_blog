package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetApi(r *gin.Engine) {
	r.POST("/sample", func(c *gin.Context) {
		c.JSON(http.StatusOK, "{hello:hello}")
	})
	r.LoadHTMLGlob("public/*")
}
