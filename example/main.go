package main

import (
	"github.com/cit965/ctp"
)

func main() {
	engine := ctp.Default()
	engine.GET("/foo", FooControllerHandler)
	g := engine.Group("/boo")
	{
		g.GET("/hello", FooControllerHandler)
		g.GET("/xx/:id", FooControllerHandler)
	}
	engine.Run(":8000")

}

func FooControllerHandler(c *ctp.Context) {
	c.Json(200, "success")
}
