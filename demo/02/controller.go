package main

import (
	"context"
	"fmt"
	"github.com/chenzebinm4/webframe"
	"log"
	"time"
)

func FooControllerHandler(c *webframe.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan any, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), 1*time.Second)
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// 处理实际工作
		time.Sleep(2 * time.Second)
		c.Json(200, "ok")
		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "time out")
		c.SetHasTimeout()
	}
	return nil
}
