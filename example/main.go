package main

import (
	"github.com/cit965/ctp"
	"time"
)

func main() {
	engine := ctp.Default()
	engine.GET("/foo", ctp.Recovery(), ctp.Log(), FooControllerHandler)
	g := engine.Group("/boo")
	{
		g.GET("/hello", FooControllerHandler)
		g.GET("/xx/:id", FooControllerHandler)
	}
	engine.Run(":8000")

}

func FooControllerHandler(c *ctp.Context) {
	time.Sleep(time.Second * 3)
	c.Json(200, "success")
}
