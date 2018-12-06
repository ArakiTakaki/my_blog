package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ArakiTakaki/my_blog/goserver/db"
	"github.com/ArakiTakaki/my_blog/goserver/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	content, err := ioutil.ReadFile("db/queries/users_table.sql")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

	r := gin.New()
	routes.SetApi(r)

	db.GetDB()
	// static files
	r.StaticFile("bundle.js", "./public/bundle.js")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.Run(":3000")
}
