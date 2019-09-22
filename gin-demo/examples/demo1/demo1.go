package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// User ...
type User struct {
	Username string
	Passwd   string
	Age      int
}

func main() {
	r := gin.New()

	// curl http://127.0.0.1:8080/
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!\n")
	})

	// curl http://127.0.0.1:8080/ping
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong\n")
	})

	// curl http://127.0.0.1:8080/user/ljs
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello %s\n", name)
	})

	// curl -X POST http://127.0.0.1:8080/form_post -d 'message=ljs&nick=ljs'
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"status_code": http.StatusOK,
				"status":      "ok",
			},
			"message": message,
			"nick":    nick,
		})
	})

	// curl -s -X PUT http://127.0.0.1:8080/post?id=ljs\&page=1 -d 'name=ljs&message=ljs'
	r.PUT("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s\n", id, page, name, message)
		c.JSON(http.StatusOK, gin.H{
			"status_code": http.StatusOK,
		})
	})

	// curl -d 'name=ljs' -F 'upload=@ljs' -X POST http://127.0.0.1:8080/upload
	r.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		fmt.Println(name)
		file, header, err := c.Request.FormFile("upload")
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request")
			return
		}
		filename := header.Filename
		fmt.Println(file, err, filename)
		c.String(http.StatusCreated, "upload successful")
	})

	r.POST("/multi/upload", func(c *gin.Context) {
		err := c.Request.ParseMultipartForm(200000)
		if err != nil {
			log.Fatal(err)
		}

		formdata := c.Request.MultipartForm
		files := formdata.File["upload"]
		for i := range files {
			file, err := files[i].Open()
			filename := files[i].Filename
			fmt.Println(file, err, filename)
			defer file.Close()
			if err != nil {
				log.Fatal(err)
			}
			c.String(http.StatusCreated, "upload successful")
		}
	})

	r.POST("/login", func(c *gin.Context) {
		var user User
		var err error
		contentType := c.Request.Header.Get("Content-Type")

		switch contentType {
		case "application/json":
			err = c.BindJSON(&user)
		case "application/x-www-form-urlencoded":
			err = c.BindWith(&user, binding.Form)
		}

		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"user":   user.Username,
			"passwd": user.Passwd,
			"age":    user.Age,
		})
	})

	// 自动Bind
	r.POST("/loginBind", func(c *gin.Context) {
		var user User

		err := c.Bind(&user)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"username": user.Username,
			"passwd":   user.Passwd,
			"age":      user.Age,
		})
	})

	// curl http://127.0.0.1:8080/render
	r.GET("/render", func(c *gin.Context) {
		contentType := c.DefaultQuery("content_type", "json")
		if contentType == "json" {
			c.JSON(http.StatusOK, gin.H{
				"user":   "rsj217-json",
				"passwd": "123",
			})
		} else if contentType == "xml" {
			c.XML(http.StatusOK, gin.H{
				"user":   "rsj217-xml",
				"passwd": "123",
			})
		}
	})

	//重定向
	// curl http://127.0.0.1:8080/redict/baidu
	r.GET("/redict/baidu", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://baidu.com")
	})

	// 全局中间件
	r.Use(MiddleWare())
	{
		r.GET("/middleware", func(c *gin.Context) {
			request := c.MustGet("request").(string)
			req, _ := c.Get("request")
			c.JSON(http.StatusOK, gin.H{
				"middile_request": request,
				"request":         req,
			})
		})
	}

	//单个路由中间件
	r.GET("/before", MiddleWare(), func(c *gin.Context) {
		request := c.MustGet("request").(string)
		c.JSON(http.StatusOK, gin.H{
			"middile_request": request,
		})
	})

	//鉴权中间件
	r.GET("/auth/signin", func(c *gin.Context) {
		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    "123",
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		c.String(http.StatusOK, "Login successful")
	})

	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "home"})
	})

	// Handle all requests using net/http
	http.Handle("/", r)

	r.Run("127.0.0.1:8080")
}

// MiddleWare 全局中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("request", "client_request")
		c.Next()
		fmt.Println("after middleware")
	}
}

// AuthMiddleWare 鉴权中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("session_id"); err == nil {
			value := cookie.Value
			fmt.Println(value)
			if value == "123" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}
}
