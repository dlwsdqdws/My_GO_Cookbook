# Hertz

- [Hertz](#hertz)
  - [Server Side](#server-side)
    - [Installation](#installation)
    - [Create a Server](#create-a-server)
    - [Routing](#routing)
      - [Static Route](#static-route)
      - [Route Group](#route-group)
      - [Param Route](#param-route)
      - [Wildcard Route](#wildcard-route)
    - [Parameter Binding](#parameter-binding)


## Server Side

### Installation

```go
go install github.com/cloudwego/hertz/cmd/hz@latest
```

### Create a Server

```go
// Default() or New()
h := server.Default(server.WithHostPoerts("127.0.0.1:8080"))
// two contexts
h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
      ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
})
h.Spin()
```

- Listening on Port 8080 by default.

### Routing

- Priority: **Static Route** > **Param Route** > **Wildcard Route**

#### Static Route

- Hertz provides `GET`, `POST`, `PUT`, `DELETE` and other methods. `ANY` can be used to register all HTTP Method methods. `Handle` can be used to register custom HTTP Method methods.

```go
func RegisterRoute(h *server.Hertz){
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
}
```

#### Route Group

- Hertz provides **Route Group** to support the routing grouping function, and middleware can also be registered to the routing group.

```go
v1 := h.Group("/v1"){
    v1.GET("/get", func(ctx context.Context, c *app.RequestContext) {
        c.String(consts.StatusOK, "get")
    })
    v1.POST("/post", func(ctx context.Context, c *app.RequestContext) {
        c.String(consts.StatusOK, "post")
    })
}
v2 := h.Group("/v2"){
    v2.PUT("/put", func(ctx context.Context, c *app.RequestContext) {
        c.String(consts.StatusOK, "put")
    })
    v2.DELETE("/delete", func(ctx context.Context, c *app.RequestContext) {
        c.String(consts.StatusOK, "delete")
    })
}
```
#### Param Route

- Hertz supports setting up routes with named parameters like `:name`, and named parameters only match a single path segment, like `/user/gordon` and `/user/you`, not `/user/profile` or `/user/`.

```go
h.GET("/hertz/:version", func(ctx context.Context, c *app.RequestContext) {
        version := c.Param("version")
        c.String(consts.StatusOK, "Hello %s", version)
    })
```

#### Wildcard Route

- Hertz supports setting routes with wildcard parameters like `*path`, and wildcard parameters will match everything containing such segment, like `/src/`, `/src/somefile.go`, `/src/subdir/somefile.go`.

```go
h.GET("/hertz/:version/*action", func(ctx context.Context, c *app.RequestContext) {
        version := c.Param("version")
        action := c.Param("action")
        message := version + " is " + action
        c.String(consts.StatusOK, message)
    })

h.POST("/hertz/:version/*action", func(ctx context.Context, c *app.RequestContext){
  //c.FullPath() == "/hertz/:version/*action"
  c.String(consts.StatusOK, c.FullPath())
})
```

### Parameter Binding

