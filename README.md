# My GO Guidebook

## A Tour of Go

https://go.dev/tour/welcome/1

## Basic Grammer

### Variable

1. `:=` and `const` can automatically infer the variable type
2. Scientific notation is supported.
   eg, `const i = 3e20 / h`

### Control Structures

#### If structure

bool expression does not need `()`, but content needs `{}`

#### Switch Structure

1. switch - case is essentially a sequence of if - else statements, that is, case can be used without constants. `break` is not needed.

```
func main() {
    fmt.Println("When's Saturday?")
    today := time.Now().Weekday()
    switch time.Saturday {
    case today + 0:
        fmt.Println("Today.")
    case today + 1:
        fmt.Println("Tomorrow.")
    case today + 2:
        fmt.Println("In two days.")
    case today + 3:
        fmt.Println("In three days.")
    case today + 4:
        fmt.Println("In four days.")
    default:
        fmt.Println("Too far away.")
    }
}
```

2. Can be used to beautify an if - else sequence

```
t := time.Now()
switch {
case t.Hour() < 12:
	fmt.Println("morning")
case t.Hour() < 18:
	fmt.Println("afternoon")
default:
	fmt.Println("evening")
}
```

#### For Structure

1. Dead loop

```
for {	}
```

2. Can be used as `while` loop in C/C++

```
i := 1
for i <= 3 {
	fmt.Println(i)
	i = i + 1
}
```
