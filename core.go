package webframe

import (
	"log"
	"net/http"
	"strings"
)

// Core 核心结构
type Core struct {
	router map[string]*Tree
}

// NewCore 初始化 Core 结构
func NewCore() *Core {
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()

	return &Core{router: router}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(url string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(url string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// FindRouteByRequest 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	// uri 和 method 全部转换为大写，保证大小写不敏感
	upperMethod := strings.ToUpper(request.Method)
	upperURI := strings.ToUpper(request.URL.Path)

	// 查找
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(upperURI)
	}

	return nil
}

// Group 初始化 Group
func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// 使用自定义的 Context
	ctx := NewContext(request, response)

	// 寻找路由
	router := c.FindRouteByRequest(request)
	if router == nil {
		_ = ctx.Json(404, "not found")
		return
	}

	if err := router(ctx); err != nil {
		_ = ctx.Json(500, "internal error")
	}

	return
}
