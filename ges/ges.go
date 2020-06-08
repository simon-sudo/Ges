package ges

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines the request handler used by ges
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 实现 serveHTTP 接口
type Engine struct {
	router map[string]HandlerFunc
}

//New ges.Engine的构造函数（constructor)
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 添加到路由表
func (engine Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

//GET 定义了添加get请求的方法
func (engine Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRouter("GET", pattern, handler)
}

//POST 定义了添加post的方法
func (engine Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRouter("POST", pattern, handler)
}

//RUN 定义启动http服务的方法
func (engine Engine) RUN(addr string) (err error) {
	//cannot use engine (variable of type Engine) as http.Handler value in argument to http.ListenAndServe: missing method ServeHTTP
	return http.ListenAndServe(addr, engine)
}

func (engine Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404: %s\n", req.URL)
	}

}
