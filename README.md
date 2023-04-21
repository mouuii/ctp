# Ctp Web Framework

Ctp is a web framework written in [Go](https://go.dev/). Now , the purpose  for this project is only to learn how to write a go http framework !

### Example

```go

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
```
