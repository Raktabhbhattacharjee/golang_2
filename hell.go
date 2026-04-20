package main

import "fmt"

type Contact struct {
	Email string
}

func (c Contact) PrintEmail() {
	fmt.Println("Email:", c.Email)
}

// Pointer method (modifies data)
func (c *Contact) UpdateEmail(newEmail string) {
	c.Email = newEmail
}

type Employee struct {
	Name string
	Contact
}

func (e Employee) PrintName() {
	fmt.Println("Name:", e.Name)
}
func main() {
	e := Employee{
		Name: "Raktabh",
		Contact: Contact{
			Email: "old@mail.com",
		},
	}
	fmt.Println("Email:", e.Email)
	e.PrintEmail()
	e.PrintName()

	// // 🔹 Update using pointer method
	e.UpdateEmail("new@mail.com")
// 
	// fmt.Println("Updated Email:", e.Email)

}
// Embedding allows the outer struct to directly access inner struct’s fields and methods
// Embedding = shortcut access to inner struct