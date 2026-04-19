package main

import "fmt"

// Struct = collection of data (like class without methods inside)
type User struct {
    Name string
    Age  int
}

// Method with POINTER receiver (*User)
// (u *User) = receiver (like 'this' in C++ / 'self' in Python)
//
// IMPORTANT:
// - This method receives a POINTER to the original object
// - So it can MODIFY the real data
func (u *User) ChangeAgevalue(newAge int) {
    // u points to original memory → change persists
    u.Age = newAge
}

func main() {
    // Creating a struct instance
    u := User{Name: "Raktabh", Age: 20}

    fmt.Println(u) // {Raktabh 20}

    // Calling method
    // Even though method expects *User,
    // Go automatically converts:
    // u.ChangeAgevalue(23) → (&u).ChangeAgevalue(23)
    u.ChangeAgevalue(23)

    // Since pointer receiver was used,
    // original struct is modified
    fmt.Println("After Value Method:", u.Age) // 23 ✅
}


/*
==========================
IMPORTANT NOTES (REVISION)
==========================

1. Go ALWAYS passes by value

2. Difference:

   func (u User)
   → copy of struct
   → changes DO NOT affect original ❌

   func (u *User)
   → copy of pointer (address)
   → both refer to SAME memory
   → changes affect original ✅

3. Go Magic:
   u.method() works for both:
   - value receiver
   - pointer receiver

   Go auto handles:
   u.method() → (&u).method() if needed

4. Golden Rule:
   If method modifies struct → use POINTER receiver

5. Mental Model:
   Value  → separate copy
   Pointer → shared memory

*/