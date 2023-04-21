package ctp

import "log"

// recovery机制，将协程中的函数异常进行捕获
func Recovery() HandlerFunc {
	// 使用函数回调
	return func(c *Context) {
		// 核心在增加这个recover机制，捕获c.Next()出现的panic
		defer func() {
			if err := recover(); err != nil {
				c.Json(500, err)
			}
		}()
		// 使用next执行具体的业务逻辑
		c.Next()
	}
}

func Log() HandlerFunc {
	return func(c *Context) {
		log.Println("-----middle------before")
		c.Next()
		log.Println("-----middle------after")
	}
}
