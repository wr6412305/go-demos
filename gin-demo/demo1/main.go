package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getApi(c *gin.Context) {
	fmt.Println(c.Query("id"))
	c.String(http.StatusOK, "ok")
}

func postApi(c *gin.Context) {
	fmt.Println(c.PostForm("id"))
	c.String(http.StatusOK, "ok")
}

func postjson(c *gin.Context) {
	var data = &struct {
		Name string `json:"title"`
	}{}

	c.BindJSON(data)

	fmt.Println(data)
	c.String(http.StatusOK, "ok")
}

func main() {
	r := gin.Default()
	r.GET("/getApi ", getApi)      //注册接口
	r.POST("/postApi ", postApi)   //注册接口
	r.POST("/postjson ", postjson) //注册接口
	r.Run()                        // listen and serve on 0.0.0.0:8080
}
