# My GO Cookbook

- [My GO Cookbook](#my-go-cookbook)
  - [Basic Grammer](#basic-grammer)
  - [Concurrent and Parallel](#concurrent-and-parallel)
  - [Dependency Management](#dependency-management)
    - [GO Path](#go-path)
    - [GO Vendor](#go-vendor)
    - [GO Module](#go-module)
  - [Test](#test)
    - [Unit Test](#unit-test)
    - [Mock](#mock)
    - [Benchmark](#benchmark)
  - [Code Optimization](#code-optimization)
    - [Go Style](#go-style)
      - [Standard](#standard)
      - [Format Tools](#format-tools)
      - [Comment](#comment)
      - [Naming Conventions](#naming-conventions)
        - [Variable](#variable)
        - [Function](#function)
        - [Package](#package)
      - [Code](#code)
        - [Control Flow](#control-flow)
        - [Error Handling](#error-handling)
    - [Performance Optimization](#performance-optimization)
      - [Slice](#slice)
      - [Map](#map)
      - [String](#string)
      - [Struct](#struct)
      - [Atomic Package](#atomic-package)
  - [Framework](#framework)
    - [ORM - Gorm](#orm---gorm)
      - [Installation](#installation)
      - [Connect to Database](#connect-to-database)
      - [CRUD](#crud)
        - [Create](#create)
        - [Read](#read)
    - [RPC - Kitex](#rpc---kitex)
    - [HTTP - Hertz](#http---hertz)
  - [Useful Tools](#useful-tools)
  - [Acknowledgements](#acknowledgements)

**WELLCOME TO EDIT**

## Basic Grammer

## Concurrent and Parallel

## Dependency Management

### GO Path

- bin : compiled binaries
- pkg : compiled intermediate products to speed up compilation
- src : source code

Disadvantage : Unable to achieve Version Control.

### GO Vendor

- vendor : place a copy of all dependent packages.

Disadvantage : Dependencies conflict.

### GO Module

- go.mod : identify module path and version information (${MAJOR}.{MINOR}.${PATCH}), describe unit dependencies (including labeling indrect and incompatible dependencies). When compiling, go will choose the lowest compatible version.
- Proxy : cache version content to achieve reliable dependency distribution.
- go get/mod : local tools

## Test

### Unit Test

- All test files end with `_test.go`
- Test function `func TestXxxx(t *testing.T)`
- Initialization should be included in `TestMain` function.

```go
func TestMain(m *testing.M) {
   // do init
   code := m.Run() // run test
   // do close and release
   os.Exit(code)
}
```

- `assert` package

```go
import "github.com/stretchr/testify/assert"

func TestFunc(t *testing.T) {
    output := Func()
    expectOutput := ""
    assert.Equal(t, expectOutput, output)
}
```

- **Code coverage** is the standard for judging whether the tests are good or bad. Using `--cover` can get the code coverage. It is good to make it more than 50%, better for more than 80%.

For more examples, please refer to https://juejin.cn/post/6908938380114034701.

### Mock

- Use the reflection mechanism to replace some functions such as R/W to achieve idempotent and make sure that the test is not affected by the environment or other users.
- [monkey package](https://github.com/bouk/monkey) : can mock methods or instances.

```go
// replace a funtion with another
func Patch(target, replacement interface{}) *PatchGuard {
	t := reflect.ValueOf(target)
	r := reflect.ValueOf(replacement)
	patchValue(t, r)
	return &PatchGuard{t, r}
}

// remove all monkey patches on target
// return if the target was patched in the first place
func Unpatch(target interface{}) bool {
	return unpatchValue(reflect.ValueOf(target))
}

func TestProcessFirstLineWithMock(t *testing.T) {
	monkey.Patch(ReadFirstLine, func() string {
		return "line110"
	})
	defer monkey.Unpatch(ReadFirstLine)
	line := ProcessFirstLine()
	assert.Equal(t, "line000", line)
}
```

### Benchmark

- Very similar with unit test. It is to test the performance of the code and the consumption of CPU.
- Start with `Benchmark` and using `gobench xxx_test.go` to run.

```go
// Serial
func BenchmarkRandom(b *testing.B) {
   InitServerIndex()
   b.ResetTimer()
   for i := 0; i < b.N; i++ {
      Random(10)
   }
}

// Parallel
func BenchmarkRandomParallel(b *testing.B) {
   InitServerIndex()
   b.ResetTimer()
   b.RunParallel(func(pb *testing.PB) {
      for pb.Next() {
         Random(10)
      }
   })
}
```

## Code Optimization

### Go Style

#### Standard

- Simplicity : eliminate unnecessary expressions
- Readability : easy to understand
- Productivity : easy to cooperate with teammates

#### Format Tools

Two useful packages.

- [gofmt](https://pkg.go.dev/cmd/gofmt)
- [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)

#### Comment

- Good code has lots of comments, bad code requires lots of comments.
- Both `/* */` and `//` are supported.
- Should explain what the code does, why it works, why it is needed and what goes wrong.
- Always add comments to the public symbols like `Read`.

```go
/*
  A Builder is used to efficiently build a string using Write methods.
  It minimizes memory copying. The zero value is ready to use.
  Do not copy a non-zero Builder.
*/
type Builder struct {
    addr *Builder // of receiver, to detect copies by value
    buf  []byte
}
```

For more exmaples please refer to https://juejin.cn/post/7189519144897740861.

#### Naming Conventions

##### Variable

- Acronyms should be all uppercase, eg `ServeHTTP`. but when it is at the beginning of the variable and does not need to be exported, could be all lowercase, eg `xmlHTTPRequest`.
- The farther a variable is from where it is used, the more contextual information its name needs to carry, eg `deadline` instead of `t`.

##### Function

- function name should be as short as possible and it does not need to carry the information of the package.
- Type information can be omitted when the return type's name is consistent with package name. eg,

```go
package http

// bad
func ServeHTTP(I net.Listener, handler Handler) error

// good
func Serve(I net.Listener, handler Handler) error
```

- Type information should be added to the function name when the package name is not consistent with the return type's name of the function.

##### Package

- Consists of lowercase letters only.
- Be short and contain some contextual information.
- Do not use the same name as the standard library.
- Better not to use commonly used variable names like `bufio` instead of `buf`.
- Better to use singular instead of plural like `encoding` instead of `encodings`.
- Use abbreviations sparingly.

#### Code

##### Control Flow

- Avoid nesting

```go
// bad
if foo {
    return x
} else {
    return nil
}
​
// good
if foo {
    return x
}
return nil
```

- Keep normal code paths with minimal indentation

```go
// bad
func OneFunc() error {
    err := doSomething()
    if err == nil{
        err := doAnotherThing()
        if err == nil{
            return nil
        }
        return err
    }
    return err
}

// good
func OneFunc() error {
    if err := doSomething(); err != nil{
        return err
    }
    if err := doAnotherThing(); err != nil{
        return err
    }
    return nil
}
```

##### Error Handling

- For simple errors that occur few times, use `errors.New()`.
- To format errors can use `fmt.Errorf`.
- For complicated errors, use Wrap and Unnwrap.

```go
list, _, err := c.GetBytes(cache.Subkey(a.actionID, "srcfiles"))
if err != nil {
    return fmt.Errorf("reading srcfiles list: %w", err)
}
```

- Use `errors.Is` instead of `==` to determine whether an error is a specific error. Use `errors.As` can get the specific kind of error.

```go
data, err = lockedfile.Read(targ)
if errors.Is(err, fs.ErrNotExist) {
    return []byte{}, nil
}else{
    // do something
}

if _, err := os.Open("non-existing"); err != nil {
    var pathError *fs.PathError
    if errors.As(err, &pathError) {
        fmt.Println("Failed at path:", pathError.Path)
    } else {
        // do something
    }
}
```

- `panic/recover` is not recommended unless it is really a huge problem. `recover` should be used in `defer` and takes effect only on the current goroutine.

Note that `defer` is actually a stack.

```go
func main(){
    defer fmt.Printf("1")
    defer fmt.Printf("2")
    defer fmt.Printf("3")
}
// output : 321
```

### Performance Optimization

#### Slice

- Better to provide size when `make` a slice.
- Better to use `copy` to replace re-slice.

#### Map

- Better to provide size when `make` a map.

#### String

- Use `strings.Builder` to replace `+` or `bytes.Buffer`.
- Better to provide size using `strings.Builder.Grow(size)`.

#### Struct

- Use an empty struct as a placeholder.

```go
m ：= make(map[int]struct{})

for i:= 0; i < n; i++{
    m[i] = struct{}{}
}
```

- `map`[key](Empty struct) = `set`

#### Atomic Package

- Atomic package provides atomic operations which can replace lock with a higher efficiency.

```go
var x int32 = 100
func f_add() {
	atomic.AddInt32(&x, 1)
}

func f_sub() {
	atomic.AddInt32(&x, -1)
}

func main() {
	for i := 0; i < 100; i++ {
		f_add()
		f_sub()
	}
	fmt.Printf("x: %v\n", x)   // 100
}
```

- For non-numeric operations, can use `atomic.Value`, which can carry an `interface{}`.

## Framework

### ORM - Gorm

#### Installation

```go
go get -u gorm.io/gorm
// take mysql as an example
go get -u gorm.io/driver/mysql
```

#### Connect to Database

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

#### CRUD

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

##### Create

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

##### Read

- `First` method returns the first data that meets the specified criteria, `ErrRecodeNotFound` if no such data.

```go
u := &Prodyct{}
db.First(u)
```

- `Find` method returns multiple data meets `where` criteria, nothing if no such data.

```go
// Get first matched record
db.Where("name = ?", "jinzhu").First(&user)
// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

// Get all matched records
db.Where("name <> ?", "jinzhu").Find(&users)
// SELECT * FROM users WHERE name <> 'jinzhu';

// IN
db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

// LIKE
db.Where("name LIKE ?", "%jin%").Find(&users)
// SELECT * FROM users WHERE name LIKE '%jin%';

// AND
db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

// Time
db.Where("updated_at > ?", lastWeek).Find(&users)
// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

// BETWEEN
db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
```

### RPC - Kitex

### HTTP - Hertz

## Useful Tools

1. [A Tour of Go](https://go.dev/tour/welcome)
2. [Effective Go](https://go.dev/doc/effective_go)

## Acknowledgements

Many thanks to Kechun Wang, Zheng Zhao, Lei Zhang from ByteDance for their help.
