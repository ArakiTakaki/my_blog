package main

import (
	"github.com/ArakiTakaki/my_blog/goserver/db"
	"github.com/ArakiTakaki/my_blog/goserver/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	database := db.GetDB()
	db.Migration(database)
	database.Close()

	server := gin.New()
	routes.SetApi(server)

	server.Run(":3000")
}
