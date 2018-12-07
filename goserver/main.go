package main

import (
	"github.com/ArakiTakaki/my_blog/goserver/db"
	"github.com/ArakiTakaki/my_blog/goserver/routes"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	database := db.GetDB()
	db.Migration(database)
	database.Close()
	server := gin.New()
	// session
	store := sessions.NewCookieStore([]byte("secret"))
	server.Use(sessions.Sessions("MySession", store))

	// routing
	routes.SetApi(server)

	server.Run(":3000")
}
