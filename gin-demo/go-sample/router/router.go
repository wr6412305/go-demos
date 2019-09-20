package router

import (
	"go-demos/gin-demo/go-sample/apis"

	"github.com/gin-gonic/gin"
)

// SetupRouter ...
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/welcome", apis.Welcome)

	user := r.Group("/user")
	{
		user.GET("/index", apis.UserIndex)
	}

	return r
}
