package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gocraft/web"
)

type AppContext struct {
	HelloCount int
}

func (c *AppContext) SetHelloCount(
	rw web.ResponseWriter,
	req *web.Request,
	next web.NextMiddlewareFunc,
) {
	c.HelloCount = 3
	next(rw, req)
}

func (c *AppContext) SayHello(
	rw web.ResponseWriter,
	req *web.Request,
) {
	fmt.Fprint(rw, strings.Repeat("Hello ", c.HelloCount), "World!")
}

func main() {
	router := web.New(AppContext{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Middleware((*AppContext).SetHelloCount).
		Get("/", (*AppContext).SayHello)
	http.ListenAndServe("localhost:3000", router)
}
