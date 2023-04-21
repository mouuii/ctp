package ctp

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type Engine struct {
	router      map[string]*Tree // all routers
	middlewares []HandlerFunc    // 中间件
}

func Default() *Engine {
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Engine{router: router}
}

func (e *Engine) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	ctx := NewContext(request, response)
	handlers := e.FindRouteByRequest(request)

	if handlers == nil {
		response.WriteHeader(404)
		log.Println("没有注册handler")
		return
	}
	log.Println("执行注册的handler")
	ctx.SetHandlers(handlers)
	ctx.Next()
}

func (e *Engine) Run(addr string) {
	log.Println("start to listen on port:", addr)
	http.ListenAndServe(addr, e)
}

func (e *Engine) Group(prefix string) *Group {
	return NewGroup(e, prefix)
}

func (e *Engine) GET(url string, handlers ...HandlerFunc) {
	if err := e.router["GET"].AddRouter(url, handlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (e *Engine) POST(url string, handlers ...HandlerFunc) {
	if err := e.router["POST"].AddRouter(url, handlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (e *Engine) PUT(url string, handlers ...HandlerFunc) {
	if err := e.router["PUT"].AddRouter(url, handlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (e *Engine) DELETE(url string, handlers ...HandlerFunc) {
	if err := e.router["DELETE"].AddRouter(url, handlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (e *Engine) FindRouteByRequest(request *http.Request) []HandlerFunc {
	// uri 和 method 全部转换为大写，保证大小写不敏感
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	// 查找第一层map
	if methodHandlers, ok := e.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

// 注册中间件
func (e *Engine) Use(middlewares ...HandlerFunc) {
	e.middlewares = append(e.middlewares, middlewares...)
}
