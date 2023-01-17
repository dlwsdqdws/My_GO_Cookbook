# My GO Cookbook

- [My GO Cookbook](#my-go-cookbook)
  - [Basic Grammer](#basic-grammer)
    - [Variable](#variable)
    - [Control Structures](#control-structures)
      - [If structure](#if-structure)
      - [Switch Structure](#switch-structure)
      - [For Structure](#for-structure)
    - [Array](#array)
    - [Slice](#slice)
    - [Map](#map)
    - [Struct](#struct)
    - [Iteration](#iteration)
      - [transverse a slice](#transverse-a-slice)
      - [transverse a map](#transverse-a-map)
      - [transverse a struct](#transverse-a-struct)
    - [Function](#function)
      - [pass by value](#pass-by-value)
      - [pass by reference](#pass-by-reference)
    - [Errors](#errors)
    - [String](#string)
      - [Format](#format)
      - [String Handling](#string-handling)
      - [Strconv](#strconv)
    - [JSON](#json)
    - [Time](#time)
    - [Processes](#processes)
  - [Parallel Programming](#parallel-programming)
  - [Useful Tools](#useful-tools)
  - [Acknowledgements](#acknowledgements)

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

1. Declaration

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

### Map

1. Declaration

```
m := make(map[key]value)
```

2. Initialization

```
var m = map[string]int{
	"four": 4,
	"five": 5,
}
```

3. CRUD

```
// create & update
m[key] = value

// read : ok == true if (key, value) in m else false
value, ok := m[key]

// delete
delete(m, key)
```

### Struct

1. Declaration

```
type user struct {
	name string
	id   int
}
```

2. Initialization

```
var user1 = user{}   // name : empty, id : 0
var user2 = user{name: "lulei"} // name : "lulei", id : 0
var user3 = user{"lulei", 1}   // name : "lulei", id : 1
```

3. Access Member
   We can access struct's member by using `.`

```
var user4 = user{}
user4.id = 2     // name : empty, id : 2
```

4. Method

```
func (u user) checkId(id int) bool {
	return u.id == id
}
```

It can be written as a function.

```
func checkIdFunc(u *user, id int) bool {
	return u.id == id
}
```

And just like function, method can also pass by reference.

### Iteration

#### transverse a slice

```
nums := []int{2, 3, 4}
for index, value := range nums {
	fmt.Println(index, value)
}
```

The order of output element traversed is determined by its index.

#### transverse a map

```
// iterate hole map
for k, v := range m {
	fmt.Println(k, v)
}

// iterate only keys
for k := range m {
	fmt.Println("key", k)
}

// iterate only values
for _, v := range m {
	fmt.Println("value", v)
}
```

The storage location of the data in the map is random, so a map can NOT be expected to return results in some desired order when traversed.

#### transverse a struct

```
// Firstly, use reflect.ValueOf() to get the reflection instance
value := reflect.ValueOf(user3)

// Secondly, traverse through NumField
for i := 0; i < value.NumField(); i++ {
    fmt.Println(i, value.Field(i))  // Thirdly, obtain the field
}
```

### Function

#### pass by value
   Function in GO is pass by value by default

```
func function_name(variable variable_type) return_type {
    return return_value
}
```

A function can return multiple values.

```
func exist(m map[string]int, k string) (v int, err bool) {
    v, err = m[k]
    return v, err
}

```

#### pass by reference
   <br>Pointers are needed to edit parameters of the function

```
func increase(a int) {
    a += 1
}

func increase2(a *int) {
    *a += 1
}

num := 1
increase(num)    // pass by value
fmt.Println(num) // 1

increase2(&num)  // pass by reference
fmt.Println(num) // 2

```

### Errors

Usually, we will return a boolean value `error` alone with return values. `nil` for no error. <br>For example, when searching for an element in an array, the return value should contain the found element and `error`. When the element is found, `error` should be `nil` otherwise it should remind the operator of an error.

```
func search(users []user, name string) (u *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("No such user")
}
```

When calling a function, we should first check whether the returned `error` reports an exception.

```
if err == nil {
	fmt.Println(u.id)
}
```

### String
#### Format

|  format   | meaning  |
|  :------:  | :-----:  |
|     %v    | return native value |
| %+v  | expand struct's names and values |
|  %#v  | value in syntax format  |
|  %b  | binary value  |
|  %f | float number |
|  %.2f | float number with 2 decimal places |

#### String Handling

|  function   | meaning  |
|  :------:  | :-----:  |
|     Join    | Concatenate string arrays (or slices) |
| Split  | Separate a string according to certain rules |
|  Count  | Counts the number of non-overlapping occurrences of a substring in a string  |
|  Fields  | Separate strings with one or more consecutive spaces  |
|  HasPrefix/HasSuffix | Check if a string has a certain prefix or suffix |
|  IndexFunc | Find index of a character or substring in a string |
|  ToLower/ToUpper | Case conversion |

For more functions like io operations, please refer to https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter02/02.1.html

#### Strconv
Strconv can realize the conversion between string and numeric types

|  function   | meaning  |
|  :------:  | :-----:  |
|     ParseInt    | Convert a string to a number, can set base and bitsize |
| ParseBool  | Convert a string to a boolean |
|  ParseFloat  | Cconvert a 'float' string to a float number |
|  ParseUnit  | Similar to the ParseInt(), but does not accept +/- signs  |
|  Atoi | Convert an integer of type string to type int |

### JSON
   

### Time
Time is a kind of data type.

|  function   | meaning  |
|  :------:  | :-----:  |
|     Now    | Current time |
| Date  | Create a Time |
|  Sub  | Time difference  |
|  Parse  | Parse a time string to a Time  |
|  Unix | Convert to unix timestamp |
|  Round | Get the Time on the hour/minute/... |

For Timer functions, please refer to https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter04/04.4.html


### Processes

## Parallel Programming


## Useful Tools

1. [A Tour of Go](https://go.dev/tour/welcome)
2. [Effective Go](https://go.dev/doc/effective_go)

## Acknowledgements

Many thanks to Kechun Wang, Zheng Zhao from ByteDance for his help.
