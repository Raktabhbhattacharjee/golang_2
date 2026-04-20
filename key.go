package main

import "fmt"

func main() {

    // ======================
    // 1. ARRAY (fixed size)
    // ======================
    var arr [2]string
    arr[0] = "hello"
    arr[1] = "Raktabh"

    fmt.Println("Array:", arr)
    fmt.Println("Array index 0:", arr[0])


    // ======================
    // 2. SLICE (dynamic)
    // ======================
    s := []string{"hello", "Raktabh"}
    fmt.Println("Slice:", s)


    // ======================
    // 3. APPEND (add element)
    // ======================
    s = append(s, "Go")
    fmt.Println("After append:", s)


    // ======================
    // 4. LENGTH & CAPACITY
    // ======================
    fmt.Println("Length:", len(s))
    fmt.Println("Capacity:", cap(s))


    // ======================
    // 5. MAKE (controlled slice)
    // ======================
    s2 := make([]int, 2, 5)
    fmt.Println("Make slice:", s2)
    fmt.Println("Len:", len(s2), "Cap:", cap(s2))


    // ======================
    // 6. SLICING
    // ======================
    nums := []int{1, 2, 3, 4}
    sub := nums[1:3]
    fmt.Println("Sliced:", sub) // [2 3]


    // ======================
    // 7. LOOP
    // ======================
    for i, v := range s {
        fmt.Println("Index:", i, "Value:", v)
    }


    // ======================
    // 8. SHARED MEMORY (IMPORTANT)
    // ======================
    a := []int{1, 2, 3}
    b := a

    b[0] = 100
    fmt.Println("Original affected:", a) // [100 2 3]


    // ======================
    // 9. COPY (safe way)
    // ======================
    c := make([]int, len(a))
    copy(c, a)

    c[0] = 999
    fmt.Println("Copy safe:", c)
    fmt.Println("Original:", a)


    // ======================
    // 10. REMOVE ELEMENT
    // ======================
    remove := []int{1, 2, 3, 4}
    i := 1
    remove = append(remove[:i], remove[i+1:]...)
    fmt.Println("After remove:", remove) // [1 3 4]


    // ======================
    // 11. STACK (LIFO)
    // ======================
    stack := []int{}
    stack = append(stack, 1)
    stack = append(stack, 2)

    stack = stack[:len(stack)-1]
    fmt.Println("Stack:", stack)


    // ======================
    // 12. QUEUE (FIFO)
    // ======================
    queue := []int{}
    queue = append(queue, 1)
    queue = append(queue, 2)

    queue = queue[1:]
    fmt.Println("Queue:", queue)
}