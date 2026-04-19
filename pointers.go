package main

import "fmt"

func main() {
    // 1. normal variable
    a := 10

    // 2. pointer to a
    p := &a

    // 3. print value and address
    fmt.Println("a:", a)
    fmt.Println("address of a:", p)

    // 4. dereference pointer
    fmt.Println("value via pointer:", *p)

    // 5. modify using pointer
    *p = 25

    // 6. check updated value
    fmt.Println("updated a:", a)
}