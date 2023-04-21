package ctp

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type Engine struct {
	router map[string]HandlerFunc
}

func Default() *Engine {
	return &Engine{router: map[string]HandlerFunc{}}
}

func (e *Engine) GET(url string, handler HandlerFunc) {
	e.router[url] = handler
}

func (e *Engine) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	ctx := NewContext(request, response)
	path := strings.Trim(request.URL.Path, "/")
	router, ok := e.router[path]
	if !ok {
		response.WriteHeader(404)
		log.Println("没有注册handler")
		return
	}
	log.Println("执行注册的handler")

	router(ctx)
}

func (e *Engine) Run(addr string) {
	log.Println("start to listen on port:", addr)
	http.ListenAndServe(addr, e)
}
