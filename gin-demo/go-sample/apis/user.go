package apis

import (
	"net/http"

	"go-demos/gin-demo/go-sample/models"

	"github.com/gin-gonic/gin"
)

// UserIndex ...
// curl http://127.0.0.1:8080/user/index
func UserIndex(c *gin.Context) {
	var user models.User
	result := models.DB.Take(&user).Value
	// fmt.Printf("result: %+v", result)

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}
