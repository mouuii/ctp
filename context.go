package ctp

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Context is the most important part of ctp. It allows us get http request and response
type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		Request: r,
		Writer:  w,
	}
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.Request.Context().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.Request.Context().Done()
}

func (ctx *Context) Err() error {
	return ctx.Request.Context().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.Request.Context().Value(key)
}

func (ctx *Context) Json(status int, obj interface{}) {

	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.Writer.WriteHeader(500)
		log.Print(err)
	}
	ctx.Writer.Write(byt)
}
