package users

import (
	"fmt"

	"gorm.io/gorm"
)

// User reprents the actor of the system. It is one of the domain of the system
type User struct {
	// Unique identifier to identify the user
	ID uint `bson:"id" json:"id"`

	// First name of the user
	FirstName string `bson:"firstName" json:"firstName"`

	// Last name of the user
	LastName string `bson:"lastName" json:"lastName"`

	// Email id of the user
	Email string `bson:"email" json:"email"`
}

// NewUser is the helper method to create a bare minimum user instance
func NewUser(id uint, firstName, lastName, email string) *User {
	return &User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}

// Implement the `stringer` interface for `User` type
func (u User) String() string {
	return fmt.Sprintf("'[%v] %s %s %s'", u.ID, u.FirstName, u.LastName, u.Email)
}

// MySqlUser is the specialized variant of the user which represents a record of user, inside MySQL database.
// It represents the User data model.
type MySqlUser struct {
	User
	gorm.Model
}
