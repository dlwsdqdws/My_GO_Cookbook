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


# Code Optimization

## Go Style

### Standard

- Simplicity : eliminate unnecessary expressions
- Readability : easy to understand
- Productivity : easy to cooperate with teammates

### Format Tools

Two useful packages.

- [gofmt](https://pkg.go.dev/cmd/gofmt)
- [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)

### Comment

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

### Naming Conventions

#### Variable

- Acronyms should be all uppercase, eg `ServeHTTP`. but when it is at the beginning of the variable and does not need to be exported, could be all lowercase, eg `xmlHTTPRequest`.
- The farther a variable is from where it is used, the more contextual information its name needs to carry, eg `deadline` instead of `t`.

#### Function

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

#### Package

- Consists of lowercase letters only.
- Be short and contain some contextual information.
- Do not use the same name as the standard library.
- Better not to use commonly used variable names like `bufio` instead of `buf`.
- Better to use singular instead of plural like `encoding` instead of `encodings`.
- Use abbreviations sparingly.

### Code

#### Control Flow

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

#### Error Handling

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

## Performance Optimization

### Slice

- Better to provide size when `make` a slice.
- Better to use `copy` to replace re-slice.

### Map

- Better to provide size when `make` a map.

### String

- Use `strings.Builder` to replace `+` or `bytes.Buffer`.
- Better to provide size using `strings.Builder.Grow(size)`.

### Struct

- Use an empty struct as a placeholder.

```go
m ：= make(map[int]struct{})

for i:= 0; i < n; i++{
    m[i] = struct{}{}
}
```

- `map`[key](Empty struct) = `set`

### Atomic Package

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