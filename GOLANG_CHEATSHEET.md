# Go (Golang) Cheatsheet

## Table of Contents

1. [Basics](#basics)
2. [Variables & Types](#variables--types)
3. [Pointers](#pointers)
4. [Structs](#structs)
5. [Functions](#functions)
6. [Control Flow](#control-flow)
7. [Arrays & Slices](#arrays--slices)
8. [Maps](#maps)
9. [Loops](#loops)
10. [Interfaces](#interfaces)
11. [Methods](#methods)
12. [Error Handling](#error-handling)
13. [Packages & Imports](#packages--imports)
14. [Concurrency](#concurrency)
15. [Common Functions](#common-functions)

---

## Basics

### Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Comments

```go
// Single line comment

/* Multi-line
   comment */
```

---

## Variables & Types

### Variable Declaration

```go
// Using var
var x int = 10
var y string = "hello"

// Type inference
var z = 42  // int
var name = "Go"  // string

// Short declaration (inside functions only)
x := 10
y := "hello"

// Multiple variables
var a, b, c int = 1, 2, 3
x, y := 10, 20
```

### Data Types

```go
// Integers
var a int8 = 127        // -128 to 127
var b int16 = 32767     // -32,768 to 32,767
var c int32 = 2147483647
var d int64 = 9223372036854775807
var e int = 0           // platform-dependent (32 or 64 bit)

// Unsigned Integers
var ua uint8 = 255      // 0 to 255
var ub uint16 = 65535
var uc uint32 = 4294967295
var ud uint64 = 18446744073709551615
var ue uint = 0

// Floating Point
var f1 float32 = 3.14
var f2 float64 = 2.718281828

// Boolean
var flag bool = true
var done bool = false

// String
var message string = "Hello"

// Rune (Unicode character)
var char rune = 'A'

// Byte (unsigned 8-bit)
var b byte = 65
```

### Constants

```go
const PI = 3.14159
const NAME = "John"

// Multiple constants
const (
    SPRING = 1
    SUMMER = 2
    FALL = 3
    WINTER = 4
)

// Iota (auto-increment)
const (
    FIRST = iota   // 0
    SECOND         // 1
    THIRD          // 2
)
```

### Type Conversion

```go
var i int = 42
var f float64 = float64(i)
var s string = "123"
num, _ := strconv.Atoi(s)  // String to int
str := strconv.Itoa(42)    // Int to string
```

---

## Pointers

### Pointer Basics

```go
var p *int          // Pointer to int
x := 10
p = &x              // Address of operator
value := *p         // Dereference operator
```

### Pointer Examples

```go
func main() {
    name := "Alice"
    ptr := &name

    fmt.Println(*ptr)  // Alice
    *ptr = "Bob"
    fmt.Println(name)  // Bob
}
```

### Nil Pointers

```go
var p *int
if p == nil {
    fmt.Println("Pointer is nil")
}
```

---

## Structs

### Struct Declaration

```go
type Person struct {
    Name string
    Age  int
    Email string
}
```

### Creating Structs

```go
// Method 1: Positional arguments
p1 := Person{"Alice", 30, "alice@example.com"}

// Method 2: Named fields
p2 := Person{
    Name: "Bob",
    Age: 25,
    Email: "bob@example.com",
}

// Method 3: Partial initialization
p3 := Person{Name: "Charlie"}

// Using new
p4 := new(Person)
p4.Name = "Diana"
p4.Age = 28
```

### Nested Structs

```go
type Address struct {
    Street string
    City   string
}

type Employee struct {
    Name    string
    Address Address
}

emp := Employee{
    Name: "John",
    Address: Address{
        Street: "123 Main St",
        City: "New York",
    },
}

fmt.Println(emp.Address.City)  // New York
```

### Embedded Structs (Composition)

```go
type Person struct {
    Name string
}

type Developer struct {
    Person  // Embedded
    Language string
}

dev := Developer{
    Person: Person{Name: "Alice"},
    Language: "Go",
}

fmt.Println(dev.Name)      // Can access embedded field directly
```

---

## Functions

### Function Declaration

```go
func add(a int, b int) int {
    return a + b
}

// Shorthand parameters of same type
func multiply(x, y int) int {
    return x * y
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Named return values
func calc(x, y int) (sum int, product int) {
    sum = x + y
    product = x * y
    return
}
```

### Variadic Functions

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

result := sum(1, 2, 3, 4, 5)  // 15
```

### Anonymous Functions & Closures

```go
// Anonymous function
func() {
    fmt.Println("Anonymous function")
}()

// Closure
counter := func() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}()

fmt.Println(counter())  // 1
fmt.Println(counter())  // 2
```

### Function as Variable

```go
var greet func(string)
greet = func(name string) {
    fmt.Println("Hello, " + name)
}

greet("Alice")
```

### Defer

```go
func main() {
    defer fmt.Println("Third")
    defer fmt.Println("Second")
    fmt.Println("First")
}
// Output: First, Second, Third
```

---

## Control Flow

### If/Else

```go
if x > 10 {
    fmt.Println("x is greater than 10")
} else if x == 10 {
    fmt.Println("x is 10")
} else {
    fmt.Println("x is less than 10")
}

// If with short statement
if x := 20; x > 10 {
    fmt.Println(x)
}
```

### Switch

```go
switch day := "Monday"; day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("End of week")
default:
    fmt.Println("Midweek")
}

// Switch with multiple cases
switch num := 2; num {
case 1, 2, 3:
    fmt.Println("Small number")
case 4, 5, 6:
    fmt.Println("Medium number")
}

// Switch without condition (like if/else)
switch {
case x < 0:
    fmt.Println("Negative")
case x == 0:
    fmt.Println("Zero")
case x > 0:
    fmt.Println("Positive")
}
```

---

## Arrays & Slices

### Arrays

```go
// Fixed size array
var arr [5]int
arr = [5]int{1, 2, 3, 4, 5}

// Array with initialization
arr := [3]string{"a", "b", "c"}

// Length of array
len(arr)

// Iteration
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}
```

### Slices

```go
// Slice (dynamic array)
var slice []int
slice = []int{1, 2, 3}

// Slice from array
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]  // [2, 3, 4]
slice := arr[:3]   // [1, 2, 3]
slice := arr[2:]   // [3, 4, 5]

// Using make
slice := make([]int, 5)        // length 5, capacity 5
slice := make([]int, 5, 10)    // length 5, capacity 10

// Append
slice = append(slice, 4, 5, 6)

// Length and Capacity
len(slice)
cap(slice)

// Copy
dest := make([]int, len(source))
copy(dest, source)
```

---

## Maps

### Map Declaration & Initialization

```go
// Declare map
var m map[string]int
m = make(map[string]int)

// Initialize
m := map[string]int{"Alice": 30, "Bob": 25}

// Using make
m := make(map[string]string)
m["name"] = "John"
m["city"] = "NYC"

// Get value
age := m["Alice"]  // 30

// Check if key exists
value, exists := m["Alice"]
if exists {
    fmt.Println(value)
}

// Delete
delete(m, "Alice")

// Length
len(m)

// Iteration
for key, value := range m {
    fmt.Println(key, value)
}

// Nil map
var emptyMap map[string]int
// emptyMap == nil  // true
```

---

## Loops

### For Loop

```go
// Traditional for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// While-like loop
for i := 0; i < 10; {
    fmt.Println(i)
    i++
}

// Infinite loop (with break)
for {
    fmt.Println("Running")
    break
}

// For range
arr := []int{1, 2, 3}
for index, value := range arr {
    fmt.Println(index, value)
}

// Only value
for _, value := range arr {
    fmt.Println(value)
}

// Only index
for index := range arr {
    fmt.Println(index)
}

// Range with map
m := map[string]int{"a": 1, "b": 2}
for key, value := range m {
    fmt.Println(key, value)
}

// Range with string
for i, char := range "Hello" {
    fmt.Println(i, char)  // char is rune
}
```

### Break & Continue

```go
for i := 0; i < 10; i++ {
    if i == 2 {
        continue
    }
    if i == 8 {
        break
    }
    fmt.Println(i)
}
```

---

## Interfaces

### Interface Declaration

```go
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}
```

### Implementing Interfaces

```go
type MyType struct {
    data string
}

func (m MyType) Read(b []byte) (int, error) {
    // Implementation
    return 0, nil
}

// MyType implements Reader interface
```

### Empty Interface

```go
var i interface{}  // Can hold any type
i = 42
i = "hello"
i = 3.14
```

### Type Assertion

```go
var i interface{} = "hello"

str, ok := i.(string)
if ok {
    fmt.Println(str)
}

// Type switch
switch v := i.(type) {
case string:
    fmt.Println("String:", v)
case int:
    fmt.Println("Integer:", v)
default:
    fmt.Println("Unknown type")
}
```

---

## Methods

### Method Declaration

```go
type Circle struct {
    Radius float64
}

// Method with receiver
func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

// Method with pointer receiver (can modify)
func (c *Circle) SetRadius(r float64) {
    c.Radius = r
}

// Usage
circle := Circle{Radius: 5}
fmt.Println(circle.Area())
circle.SetRadius(10)
```

---

## Error Handling

### Error Interface

```go
type error interface {
    Error() string
}
```

### Creating Errors

```go
import "errors"

// Using errors.New
err := errors.New("something went wrong")

// Using fmt.Errorf
err := fmt.Errorf("invalid input: %s", input)

// Custom error
type CustomError struct {
    Code    int
    Message string
}

func (e CustomError) Error() string {
    return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}
```

### Error Handling

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Handling error
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println(result)
```

### Panic & Recover

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from:", r)
    }
}()

panic("Something went wrong!")
```

---

## Packages & Imports

### Imports

```go
import (
    "fmt"
    "math"
    "strings"
)

// Alias
import (
    f "fmt"
    m "math"
)

f.Println("Hello")
m.Sqrt(16)
```

### Exported/Unexported

```go
// Exported (starts with uppercase)
func PublicFunction() {}
var PublicVariable = 42

// Unexported (starts with lowercase)
func privateFunction() {}
var privateVariable = 42
```

### Creating a Package

```go
// File: mypackage/helper.go
package mypackage

func Helper() string {
    return "I'm a helper"
}
```

---

## Concurrency

### Goroutines

```go
go func() {
    fmt.Println("Running in goroutine")
}()

// Named function
func worker() {
    fmt.Println("Worker")
}

go worker()
```

### Channels

```go
// Create channel
ch := make(chan int)

// Send value
go func() {
    ch <- 42
}()

// Receive value
value := <-ch

// Buffered channel
ch := make(chan int, 2)
ch <- 1
ch <- 2
// ch <- 3  // Would block
```

### Select

```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() {
    ch1 <- "Hello"
}()

go func() {
    ch2 <- "World"
}()

for i := 0; i < 2; i++ {
    select {
    case msg1 := <-ch1:
        fmt.Println("Received:", msg1)
    case msg2 := <-ch2:
        fmt.Println("Received:", msg2)
    }
}
```

### WaitGroup

```go
var wg sync.WaitGroup

wg.Add(2)

go func() {
    defer wg.Done()
    fmt.Println("Task 1")
}()

go func() {
    defer wg.Done()
    fmt.Println("Task 2")
}()

wg.Wait()
```

### Mutex

```go
var mu sync.Mutex
var count int

mu.Lock()
count++
mu.Unlock()

// Better: Defer unlock
mu.Lock()
defer mu.Unlock()
count++
```

---

## Common Functions

### String Functions

```go
import "strings"

str := "Hello World"

strings.ToLower(str)           // "hello world"
strings.ToUpper(str)           // "HELLO WORLD"
strings.Contains(str, "World") // true
strings.Index(str, "World")    // 6
strings.Replace(str, "World", "Go", 1)  // "Hello Go"
strings.Split(str, " ")        // ["Hello", "World"]
strings.Join([]string{"a", "b"}, "-")  // "a-b"
strings.TrimSpace("  hello  ") // "hello"
```

### Math Functions

```go
import "math"

math.Abs(-5)            // 5
math.Sqrt(16)           // 4
math.Pow(2, 3)          // 8
math.Max(5, 3)          // 5
math.Min(5, 3)          // 3
math.Floor(3.7)         // 3
math.Ceil(3.2)          // 4
math.Round(3.5)         // 4
```

### Formatting

```go
import "fmt"

fmt.Println("Hello")
fmt.Printf("Formatted: %d\n", 42)
fmt.Sprintf("Result: %s", "Success")

// Common format specifiers
%d  - integer
%f  - float
%s  - string
%v  - any value
%T  - type
%q  - quoted string
%x  - hexadecimal
```

### Time Functions

```go
import "time"

now := time.Now()
formatted := now.Format("2006-01-02 15:04:05")
time.Sleep(time.Second)
time.After(5 * time.Second)  // Channel that fires after duration
```

---

## Quick Reference

| Operator | Description       |
| -------- | ----------------- |
| `+`      | Addition          |
| `-`      | Subtraction       |
| `*`      | Multiplication    |
| `/`      | Division          |
| `%`      | Modulus           |
| `==`     | Equal             |
| `!=`     | Not Equal         |
| `<`      | Less Than         |
| `>`      | Greater Than      |
| `<=`     | Less or Equal     |
| `>=`     | Greater or Equal  |
| `&&`     | Logical AND       |
| `\|\|`   | Logical OR        |
| `!`      | Logical NOT       |
| `:=`     | Short declaration |
| `&`      | Address of        |
| `*`      | Dereference       |

---

## Tips

1. **Error Handling**: Always check for errors explicitly
2. **Defer**: Use defer for cleanup operations
3. **Interfaces**: Write interfaces small and focused
4. **Goroutines**: Are lightweight, can run thousands safely
5. **Channels**: Prefer for communication between goroutines
6. **Zero Values**: All types have zero values (0, "", nil, false, etc.)
7. **Blank Identifier**: Use `_` to ignore unused values
8. **Testing**: Use `_test.go` files and `testing` package
