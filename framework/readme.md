# Framework

- [Framework](#framework)
  - [ORM - Gorm](#orm---gorm)
    - [Object–relational Mapping](#objectrelational-mapping)
    - [Installation](#installation)
    - [Connect to Database](#connect-to-database)
    - [CRUD](#crud)
      - [Create](#create)
      - [Read](#read)
      - [Update](#update)
      - [Delete](#delete)
    - [Transaction](#transaction)
    - [Hook](#hook)
    - [Plugins](#plugins)
  - [RPC - Kitex](#rpc---kitex)
    - [Remote Procedure Call](#remote-procedure-call)
    - [Server Side](#server-side)
      - [Installation](#installation-1)
      - [IDL](#idl)
      - [echo](#echo)
      - [handler](#handler)
    - [Client Side](#client-side)
      - [Create a Client](#create-a-client)
      - [Send a Request](#send-a-request)
    - [Service Registry and Discovery](#service-registry-and-discovery)
      - [Service Registry](#service-registry)
      - [Service Discovery](#service-discovery)
    - [Plugins](#plugins-1)
  - [HTTP - Hertz](#http---hertz)

## ORM - Gorm

### Object–relational Mapping

- Object–relational mapping (ORM) is a programming technique for converting data between a relational database and the heap of an object-oriented programming language (Wikipedia), like Mybatis in java.
- Using ORM, we can associate a data table in database with a certain class/struct, and by modifying the class/struct instance, we can easily CRUD the database without using SQL statements.

### Installation

```go
go get -u gorm.io/gorm
// take mysql as an example
go get -u gorm.io/driver/mysql
```

### Connect to Database

- Gorm can support MySQL, PostgreSQL, SQlite, SQL Server. Take SQLServer as an example.

```go
// import driver
import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
```

- DSN(data source name).

```go
"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
```

- Note that always check err before operating CRUD. Here `panic` is suggested if the database cannot be connected.

### CRUD

- GORM uses `ID` as the primary key, the snake-case of the structure name as the table name, the snake-case field name as the column name, and uses the `CreatedAt`, `UpdatedAt` fields to track creation and update time. So, `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` will be automatically created and work as their name.

```go
type Product struct {
    ID      uint    `gorm:"primarykey"`
    Code    string  `gorm:"column: code"`
    Price   uint    `gorm:"column: user_id"`

    // can set default values
    Name    string `gorm:"default:galeone"`
    Age     int64  `gorm:"default:18"`

    // the belows can be created automatically
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt sql.NullTime `gorm:"index"`
}
```

#### Create

- One piece of data.

```go
p := &Product({ Code : "042", Price: 100})
res := db.Create(p)
if res.Error != nil{
    // error handler
}
```

- Multiple pieces of data

```go
// Create a list of struct
products := []*Product{{Code : "041"}, {Code : "042"}, {Code : "043"}}
res := db.Create(products)
if res.Error != nil{
    // error handler
}
```

- Is it no need to set values for `ID`, `CreatedAt`, etc.
- Use `clause.Onconfict` to handle conflict. We cannot use `when` after `Create()`.

```go
p := &Product({ Code : "042", ID: 1})
// here we do nothing when conflict happens
db.Clauses(clause.Onconfict{DoNothing : true}).Create(&p)
```

#### Read

- `First` method returns the first data that meets the specified criteria, `ErrRecodeNotFound` if no such data.

```go
u := &Prodyct{}
db.First(u)
```

- `Find` method returns multiple data meets `where` criteria, nothing if no such data.

```go
p := &Product{}
// Get first matched record
res := db.Where("name = ?", "jinzhu").First(p)
// SELECT * FROM products WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

products := make([]*Product, 0)
// Get all matched records
res = db.Where("name <> ?", "jinzhu").Find(&products)
// SELECT * FROM products WHERE name <> 'jinzhu';

// Other Inquiries
// IN
db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&products)
// SELECT * FROM products WHERE name IN ('jinzhu','jinzhu 2');

// LIKE
db.Where("name LIKE ?", "%jin%").Find(&products)
// SELECT * FROM products WHERE name LIKE '%jin%';

// AND
db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&products)
// SELECT * FROM products WHERE name = 'jinzhu' AND age >= 22;

// Time
db.Where("updated_at > ?", lastWeek).Find(&products)
// SELECT * FROM products WHERE updated_at > '2000-01-01 00:00:00';

// BETWEEN
db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&products)
// SELECT * FROM products WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
```

- When using a struct as an inquiry, the zero values(eg, 0, false) will not be used. If zero values is needed, we can use `map` as a inquery.

```go
db.Where(map[string]interface{}{"name": "jinzhu", "age": 0}).Find(&products)
// SELECT * FROM products WHERE name = "jinzhu" AND age = 0;
```

- Slice can used as a inquery.

```go
db.Where([]int64{20, 21, 22}).Find(&products)
// SELECT * FROM products WHERE id IN (20, 21, 22);
```

#### Update

- Single

```go
db.Model(&product).Where("name = ?", "jinzhu").Update("Price", 200)
```

- Multiple

```go
db.Model(&Product{ID : 111}).Updates(Product{Name : "hello", Age : 20})
```

- Can use `map` or `Select` update zero values.

```go
db.Model(&product).Updates(map[string]interface{}{"Price": 200, "activated": false})
```

- Selected Column

```go
// only update price even though multiple columns in the map
db.Model(&Product{ID : 111}).Select("Price").Updates(map[string]interface{}{"Price": 200, "activated": false})
```

- SQL

```go
db.Model(&Product{ID : 111}).Updates("age", gorm.Expr("age * ? + ?", 2, 100))
```

#### Delete

- Hard Delete

```go
db.Delete(&p)
// DELETE from products where id = 10;

// where could be used
db.Where("name = ?", "jinzhu").Delete(&p)
// DELETE from products where id = 10 AND name = "jinzhu";

db.Delete(&User{}, 10)
// DELETE FROM products WHERE id = 10;

db.Delete(&User{}, "10")
// DELETE FROM products WHERE id = 10;

db.Delete(&products, []int{1,2,3})
// DELETE FROM products WHERE id IN (1,2,3);

db.Where("product LIKE ?", "%jinzhu%").Delete(&Product{})
// DELETE from products where product LIKE "%jinzhu%";

db.Delete(&Product{}, "product LIKE ?", "%jinzhu%")
// DELETE from products where product LIKE "%jinzhu%";
```

- Soft Delete

```go
// use gorm.DeletedAt
// when call Delete(), the data will not be deleted physically
// but label DeletedAt as current time
// when calling Find(), soft-deleted data will be ignored
type User struct {
  ID int64
  Name string
  Age int64
  Deleted gorm.DeletedAt
}

db.Where("age = ?", 20).Delete(&User{})
// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

// Use Unscoped can find soft-deleted data or realize hard delete
db.Unscoped().Where("age = 20").Find(&users)
// SELECT * FROM users WHERE age = 20;
db.Unscoped().Delete(&order)
// DELETE FROM orders WHERE id=10;
```

### Transaction

- Transaction is a sequence of database operations that are either all executed or none of them are executed. A transaction consists of all database operations performed between the start of the transaction and the end of the transaction.
- If CUD is not needed, better to disable transaction to improve performance. Use `PrepareStmt` caches prepared statements can improve the speed of subsequent calls.

```go
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  SkipDefaultTransaction: true,
  PrepareStmt: true
})
```

- Gorm Transaction provides `Begin()`, `Commit` and `Rollback()`.

```go
// use tx instead of db to start a transaction
tx := db.Begin()
​
// db operations
tx.Create(...)
​
// ...
​
// rollback if error occurs
if err := tx.Create(...).Error; err != nil{
  tx.Rollback()
  return
}
​
// otherwise submit this transaction
tx.Commit()
```

- Gorm also provides a Transaction method.

```go
if err := db.Transaction(func(tx *gorm.DB) error {
  if err := tx.Create(&User{Name: "name"}).Error; err != nil {
    // rollback automatically
    return err
  }

  if err := tx.Create(&User{Name: "name1"}).Error; err != nil {
    return err
  }

  return nil
}); err != nil{
  return
}
```

### Hook

- Hook are functions called before or after CRUD operations.

```go
type User struct {
    ID      int64
    Name    string `gorm:"default:galeone"`
    Age     string `gorm:"default:18"`
}
​
type Email struct {
    ID      int64
    Name    string
    Email   string
}
​
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    if u.Age < 0 {
        return errors.New("can't save invalid data")
    }
    return nil;
}
​
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
    return tx.Create(&Email{ID: u.ID, Email: i.Name + "@***.com"}).Error
}
```

- Hook will be called automatically when calling CRUD. If it returns an error, GORM will stop subsequent operations and rollback the transaction.

### Plugins

|            plugins            |                 links                 |
| :---------------------------: | :-----------------------------------: |
|        Code Generation        |      www.github.com/go-gorm/gen       |
| Optimizer/Index/Comment Hints |     www.github.com/go-gorm/hints      |
|        Sharding Tables        |    www.github.com/go-gorm/sharding    |
|        Optimistic Lock        | www.github.com/go-gorm/optimisticlock |
|     Read/Write Splitting      |   www.github.com/go-gorm/dbresolver   |
|         OpenTelemetry         | www.github.com/go-gorm/opentelemetry  |

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
