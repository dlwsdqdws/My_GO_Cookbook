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

1. bool expression does not need `()`, but content needs `{}`
2. Variable can be declared just behind `if`.

```
if v := math.Pow(x,y); v < limit {
    return v
}
```

#### Switch Structure

1. switch - case is essentially a sequence of if - else statements, that is, case can be used without constants. `break` is not needed.

```
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
```

2. Can be used to beautify an if - else sequence. Variable can be declared just behind switch.

```
switch t := time.Now() {
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

### Array

Declaration

```
// without initialization
var a [5]int

// with initialization
b := [5]int{1, 2, 3, 4, 5}
```

Other operations are very similar with those in C/C++.

### Slice

1. Slice itself does not store the data. It just like reference of variable-length array in C/C++.Changing a value in a slice will change the data it 'points' the array, so other slices 'point' the array(eg, copy) will also change. 
2. Declaration 
```
// part of array
a := [5]int{1, 2, 3, 4, 5}
var b []int = a[2:4]

// use make
s := make([]string, 3)

// [](type) without a number in []
good := []string{"g", "o", "o", "d"}
```
3. append operation needs to be assigned to the original slice.
```
s = append(s, "good")
```
4. Initialization
```
board := [][]int{
	[]int{1, 0, 1},
	[]int{0, 1, 1},
	[]int{1, 1, 0},
}
```

## Acknowledgements

Many thanks to Kechun Wang from ByteDance for his help.
