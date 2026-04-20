package main

import "fmt"

// Base struct
type Address struct {
    City  string
    State string
}

// Method on Address
func (a Address) PrintAddress() {
    fmt.Println("Address:", a.City, a.State)
}

// Pointer method on Address (modifies data)
func (a *Address) UpdateCity(newCity string) {
    a.City = newCity
}

// Struct with embedding
type User struct {
    Name string
    Address   // 👈 embedded
}

// Method on User
func (u User) PrintUser() {
    fmt.Println("User:", u.Name)
}

func main() {
    u := User{
        Name: "Raktabh",
        Address: Address{
            City:  "Guwahati",
            State: "Assam",
        },
    }

    // 🔥 Field promotion
    fmt.Println("City:", u.City)
    fmt.Println("State:", u.State)

    // 🔥 Method from Address (promoted)
    u.PrintAddress()

    // 🔥 Method from User
    u.PrintUser()

    // 🔥 Pointer method (modifies original)
    u.UpdateCity("Delhi")

    fmt.Println("After Update:", u.City)

    // Explicit call (same thing)
    u.Address.PrintAddress()
}