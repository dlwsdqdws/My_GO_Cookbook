# Framework

- [Framework](#framework)
  - [ORM - Gorm](#orm---gorm)
  - [RPC - Kitex](#rpc---kitex)
    - [Remote Procedure Call](#remote-procedure-call)
    - [Server Side](#server-side)
      - [Installation](#installation)
      - [IDL](#idl)
      - [echo](#echo)
      - [handler](#handler)
    - [Client Side](#client-side)
      - [Create a Client](#create-a-client)
      - [Send a Request](#send-a-request)
    - [Service Registry and Discovery](#service-registry-and-discovery)
      - [Service Registry](#service-registry)
      - [Service Discovery](#service-discovery)
    - [Plugins](#plugins)
  - [HTTP - Hertz](#http---hertz)
    - [Server Side](#server-side-1)
      - [Installation](#installation-1)
      - [Create a Server](#create-a-server)
      - [Routing](#routing)
        - [Static Route](#static-route)
        - [Route Group](#route-group)
        - [Param Route](#param-route)
        - [Wildcard Route](#wildcard-route)
      - [Parameter Binding](#parameter-binding)

## ORM - Gorm

Turn to [GORM](/framework/GORM/)

## RPC - Kitex

### Remote Procedure Call

- Remote Procedure Call(RPC) is a software communication protocol that one program can use to request a service from a program located in another computer on a network without having to understand the network's details.

### Server Side

#### Installation

```go
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
go install github.com/cloudwego/thriftgo@latest
```

#### IDL

- Interface definition language (IDL) allows a program or object written in one language to communicate with another program written in an unknown language. We can use IDL to support RPC's message transimit definition.
- Kitex supports [thrift](https://thrift.apache.org/docs/idl) and [proto3](https://developers.google.com/protocol-buffers/docs/proto3) by default, and it uses the extended thrift as the underlying transport protocol.

```go
// echo.thrift
namespace go api
​
struct Request {
    1: string message
}
​
struct Resposne {
    1: string message
}
​
service Echo {
    Reponse echo(1: Request req)
}
```

#### echo

- Generate echo service code. `-module` indicates the go module name of the generated project, `-service` indicates that we want to generate a server project, `example` is the name of the service, `echo.thrift` is IDL file.

```go
kitex -module exmaple -service example echo.thrift
```

- The generated project structure is as follows, among which, `build.sh` is the build script(code -> binary file), `kitex_gen` is the generated code (including service/client code) related to the IDL content, `main.go` is the program entry, and `handler.go` is to implement the methods defined by the IDL service in this file.

```go
.
|-- build.sh
|-- echo.thrift
|-- handler.go
|-- kitex_gen
|   `-- api
|       |-- echo
|       |   |-- client.go
|       |   |-- echo.go
|       |   |-- invoker.go
|       |   `-- server.go
|       |-- echo.go
|       `-- k-echo.go
|-- main.go
`-- script
    |-- bootstrap.sh
    `-- settings.py
```

#### handler

```go
package main
​
import (
        "context"
        api "exmaple/kitex_gen/api"
)
​
// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}
​
// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
        // TODO: Your code here...
        return
}
```

- Run `sh output/bootstrap.sh` to start the server.
- Listening on Port 8888 by default. To modify the running port, open `main.go` and specify configuration parameters for the `NewServer` function. For more information, please refer to https://juejin.cn/post/7190660194014068796#heading-9.

```go
 addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
  svr := api.NewServer(new(EchoImpl), server.WithServiceAddr(addr))
```

### Client Side

#### Create a Client

```go
import "example/kitex_gen/api/echo"
import "github.com/cloudwego/kitex/client"
...
c, err := echo.NewClient("example", client.WithHostPorts("0.0.0.0:8888"))
if err != nil {
  log.Fatal(err)
}
```

- The first parameter "example" is the service name, and the second parameter is options, which are used to pass in [parameters](https://www.cloudwego.io/zh/docs/kitex/tutorials/basic-feature/).

#### Send a Request

```go
import "example/kitex_gen/api"

// create a request named req
req := &api.Request{Message: "my request"}

// context.Context is used to transmit information or control some actions of this call
// The second parameter is the request name for this call
// The third parameter is the options, https://www.cloudwego.io/zh/docs/kitex/tutorials/basic-feature
resp, err := c.Echo(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
if err != nil {
  log.Fatal(err)
}
log.Println(resp)
```

### Service Registry and Discovery

#### Service Registry

- Service Registry: service process registers its information in the registry, usually including host, port number, protocol.

```go
type HelloImpl struct{}

// implement HelloImpl function
func (h *HelloImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error){
  resp = &api.Response{
    Message : req.Message,
  }

  return
}

func main(){
  r, err := etcd.NewEtcdRegistry([]string("127.0.0.1:2379"))
  if err != nil {
    log.Fatal(err)
  }

  // init server
  server := hello.NewServer(
    new(HelloImpl),
    server.WithRegistry(r),
    server.WithServerBasicInfo(&rpcinfo.EndPointBasicInfo{
      ServiceName : "Hello",
    }))

  err = server.Run()
  if err != nil{
    log.Fatal(err)
  }
}
```

#### Service Discovery

- Service Discovery: client process initiates a query to the registry center to obtain service information.

```go
func main(){
  e,err := etcd.NewEtcdResolver([]string("127.0.0.1:2379"))
  if err!= nil{
    log.Fatal(err)
  }

  // The first parameter is the service name
  clint := hello.MustNewClient("Hello", client.WithResolver(r))

  for {
    ctx,cancel := context.WithTimeout(context.Background(), time.Second*3)
    resp,err := client.Echo(ctx, &api.Request{
      Message : "Hello"
    })

    cancel()
    if err != nil{
      log.Fatal(err)
    }

    log.Println(resp)
    time.Sleep(time.Second)
  }
}
```

- For more examples, please refer to www.github.com/kitex-contrib/registry-etcd/tree/main/example

### Plugins

|    plugins    |                      links                      |
| :-----------: | :---------------------------------------------: |
|      XDS      |        www.github.com/kitex-contrib/xds         |
| opentelemetry | www.github.com/kitex-contrib/obs-opentelemetry  |
|     ETCD      |   www.github.com/kitex-contrib/registry-etcd    |
|     Nacos     |   www.github.com/kitex-contrib/registry-nacos   |
|   Zookeeper   | www.github.com/kitex-contrib/registry-zookeeper |
|    polaris    |      www.github.com/kitex-contrib/polaris       |

## HTTP - Hertz

### Server Side

#### Installation

```go
go install github.com/cloudwego/hertz/cmd/hz@latest
```

#### Create a Server

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

#### Routing

- Priority: **Static Route** > **Param Route** > **Wildcard Route**

##### Static Route

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
##### Param Route

- Hertz supports setting up routes with named parameters like `:name`, and named parameters only match a single path segment, like `/user/gordon` and `/user/you`, not `/user/profile` or `/user/`.

```go
h.GET("/hertz/:version", func(ctx context.Context, c *app.RequestContext) {
        version := c.Param("version")
        c.String(consts.StatusOK, "Hello %s", version)
    })
```

##### Wildcard Route

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

#### Parameter Binding

