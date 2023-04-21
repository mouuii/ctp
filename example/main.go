package main

import (
	"github.com/cit965/ctp"
	"time"
)

func main() {
	r := ctp.Default()
	r.GET("/foo", ctp.Recovery(), ctp.Log(), FooControllerHandler)
	g := r.Group("/boo")
	g.Use(ctp.Log())
	{
		g.GET("/hello", FooControllerHandler)
		g.GET("/xx/:id", FooControllerHandler)
	}
	r.Run(":8000")

}

func FooControllerHandler(c *ctp.Context) {
	time.Sleep(time.Second * 3)
	c.Json(200, "success")
}
