package ges

import (
	"log"
	"net/http"
)

// HandlerFunc defines the request handler used by ges
type HandlerFunc func(*Context)

// Engine 实现 serveHTTP 接口
type Engine struct {
	router *router
}

//New ges.Engine的构造函数（constructor)
func New() *Engine {
	return &Engine{router: newRouter()}
}

// 添加到路由表
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router.addRoute(method, pattern, handler)
}

//GET 定义了添加get请求的方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

//POST 定义了添加post的方法
func (engine Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

//RUN 定义启动http服务的方法
func (engine Engine) Run(addr string) (err error) {
	//cannot use engine (variable of type Engine) as http.Handler value in argument to http.ListenAndServe: missing method ServeHTTP
	return http.ListenAndServe(addr, engine)
}

func (engine Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
