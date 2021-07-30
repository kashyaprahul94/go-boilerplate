package main

import (
	"fmt"

	"github.com/kashyaprahul94/go-boilerplate/pkg/users"
)

func main() {
	// Simulate API request to save user
	saveUser()
}

func saveUser() {
	// Extract info from HTTP Request
	userId := 1
	firstName := "John"
	lastName := "Wick"

	// Generate a new domain user
	newUser := users.NewUser(userId, firstName, lastName)

	fmt.Printf("newUser: %v\n\n", newUser)

	// Persist it to MySQL
	newUser.InsertIntoMySql()
}
