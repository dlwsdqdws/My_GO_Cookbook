package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.Data(200, "text/plain; charset=utf-8", []byte("OK"))
	})

	h.Spin()
}
