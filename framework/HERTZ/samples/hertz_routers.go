package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	// StaticFS is used to register static files
	h.StaticFS("/", &app.FS{Root: "./", GenerateIndexPages: true})

	h.GET("/get", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "get")
	})
	h.POST("/post", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "post")
	})
	h.PUT("/put", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "put")
	})
	h.DELETE("/delete", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "delete")
	})
	h.PATCH("/patch", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "patch")
	})
	h.HEAD("/head", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "head")
	})
	h.OPTIONS("/options", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "options")
	})
	h.Any("/ping_any", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "any")
	})
	h.Handle("LOAD", "/load", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "load")
	})
	h.Spin()
}
