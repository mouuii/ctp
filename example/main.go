package main

import (
	"github.com/cit965/ctp"
)

func main() {
	engine := ctp.Default()
	engine.GET("foo", FooControllerHandler)
	engine.Run(":8000")

}

func FooControllerHandler(c *ctp.Context) {
	c.Json(200, "success")
}
