# Hertz

- [Hertz](#hertz)
  - [Server Side](#server-side)
    - [Installation](#installation)
    - [Create a Server](#create-a-server)
    - [Routing](#routing)
      - [Static Route](#static-route)
        - [Methods](#methods)
        - [Route Group](#route-group)
      - [Param Route](#param-route)
      - [Wildcard Route](#wildcard-route)
    - [Parameter Binding](#parameter-binding)
    - [Middleware](#middleware)
  - [Client Side](#client-side)
  - [Plugins](#plugins)

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

- IDL can be used.

```go
namespace go hello.example

struct HelloReq{
    // binding <-> api.query
    1 : string Name(api.query = "name");
}

struct HelloResp{
    1 : string RespBody;
}

service HelloService {
    HelloResp HelloMethod(1 : HelloReq request)(api.get = "/hello");
}
```

### Routing

- Priority: **Static Route** > **Param Route** > **Wildcard Route**

#### Static Route

##### Methods

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

##### Route Group

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

- Hertz provides `Bind`, `Validate`, `BindAndValidate` methods for binding and validate parameters.

```go
type Args struct{
    Query string `query:"query"`
    QuerySlice []string `query:"q"`
    Path string `path:"path"`
    Header string `header:"header"`
    Form string `form:"form"`
    Json string `json:"json`
    // validate
    Vd int `query:"vd" vd:"$==0||$==1"`
}

func main(){
    h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

    h.POST("v:path/bind", func(ctx context.Context, c *app.RequestContext){
        var arg Args
        // pass by reference
        err := c.BindAndValidate(&arg)
        if err != nil{
            panic(err)
        }
        fmt.Println(arg)
    })

    h.Spin()
}
```

### Middleware

- Server-side middleware is a function in the HTTP request-response cycle that provides a convenient mechanism to inspect and filter HTTP requests entering the application, such as logging each request or enabling CORS.
- Middleware can be executed before or after the request passes deeper into the logic.

```go
func MyMiddleware() app.HandlerFunc {
  return func(ctx context.Context, c *app.RequestContext) {
    // pre-handle
    // ...
    c.Next(ctx) // call the next middleware(handler)
    // post-handle
    // ...
  }
}
​
func main() {
    h := server.Default(server.WithHostPort("127.0.0.1:8080"))
    h.Use(MyMiddleware())  // Global Middleware
    h.Get("/middleware",func(ctx context.Context, c *app.RequestContext) {
        c.String(consts.StatusOK, "Hello hertz!")
    })
    h.Spin()
}
```

- Can use `Abort()`, `AbortWithMsg(msg string, statusCode int)`, `AbortWithStatus(code int)` terminates the execution of the middleware call chain.

## Client Side

- Hertz provides HTTP Client to help users send HTTP requests.

```go
c, err := client.NewClient()
if err != nil {
    return
}

// send http get request
// Get parameters : context, dst, url
status, body, err := c.Get(context.Background(), nil, "https://www.example.com")
if err != nil {
    return
}
fmt.Printf("status=%v body=%v\n", status, string(body))
​
// send http post request
var postArgs protocol.Args
postArgs.Set("arg","a") // Set post args
// Post parameters : context, dst, url
status, body, err = c.Post(context.Background(), nil, "https://www.example.com", &postArgs)
if err != nil {
    return
}
fmt.Printf("status=%v body=%v\n", status, string(body))
```

- More examples please refer to www.github.com/cloudwego/hertz-examples#client

## Plugins

|    plugins    |                     links                      |
| :-----------: | :--------------------------------------------: |
|     HTTP2     |       www.github.com/hertz-contrib/http2       |
| opentelemetry | www.github.com/hertz-contrib/obs-opentelemetry |
|     i18n      |       www.github.com/hertz-contrib/i18n        |
| Reverse Proxy |   www.github.com/hertz-contrib/reverseproxy    |
|      JWT      |        www.github.com/hertz-contrib/jwt        |
|   Websocket   |     www.github.com/hertz-contrib/websocket     |

