package main

import (
	"github.com/cit965/ctp"
	"net/http"
)

func main() {
	engin := ctp.Default()
	engin.GET("foo", FooControllerHandler)
	server := &http.Server{
		Addr:    ":8000",
		Handler: engin,
	}
	server.ListenAndServe()

}

func FooControllerHandler(c *ctp.Context) {
	c.Json(200, "success")
}
