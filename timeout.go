package ctp

import (
	"context"
	"fmt"
	"log"
	"time"
)

func TimeoutMiddleware(fun HandlerFunc, d time.Duration) HandlerFunc {
	// 使用函数回调
	return func(c *Context) {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		// 执行业务逻辑前预操作：初始化超时 context
		durationCtx, cancel := context.WithTimeout(context.Background(), d)
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			// 执行具体的业务逻辑
			fun(c)

			finish <- struct{}{}
		}()
		// 执行业务逻辑后操作
		select {
		case p := <-panicChan:
			log.Println(p)
			c.Writer.WriteHeader(500)
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			log.Println("chaoshi")
			c.Writer.Write([]byte("time out"))
		}
		return
	}
}
