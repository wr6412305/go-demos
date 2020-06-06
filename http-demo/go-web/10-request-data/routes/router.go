package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter 返回一个 mux.Router 类型指针, 从而可以当作处理器使用
func NewRouter() *mux.Router {
	// 创建 mux.Router 路由器示例
	router := mux.NewRouter().StrictSlash(true)

	router.Use(loggingRequestInfo)

	// 遍历 web.go 中定义的所有 webRoutes
	for _, route := range webRoutes {
		// 将每个 web 路由应用到路由器
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

// 记录请求日志信息中间件
func loggingRequestInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 打印请求 URL 明细
		url, _ := json.Marshal(r.URL)
		log.Println(string(url))

		// log.Printf("Request URL: %s\n", r.URL)
		// log.Printf("User Agent: %s\n", r.Header.Get("User-Agent"))
		// log.Printf("Request Header: %v\n", r.Header)

		next.ServeHTTP(w, r)
	})
}
