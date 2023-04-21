# Ctp Web Framework

Ctp is a web framework written in [Go](https://go.dev/). Now , the purpose  for this project is only to learn how to write a go http framework !

### Example

```go
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
```
