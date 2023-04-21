package ctp

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type Engin struct {
	router map[string]HandlerFunc
}

func Default() *Engin {
	return &Engin{router: map[string]HandlerFunc{}}
}

func (e *Engin) GET(url string, hander HandlerFunc) {
	e.router[url] = hander
}

func (e *Engin) ServeHTTP(response http.ResponseWriter, request *http.Request) {

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
