package middleware

import (
	"github.com/gin-gonic/gin"
)

// Cors クロスオリジン
func Cors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3030")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Accept, Content-Type")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Next()
}
