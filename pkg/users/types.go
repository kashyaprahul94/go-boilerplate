package users

import "fmt"

// User is the domain model
type User struct {
	Id        int
	FirstName string
	LastName  string
}

func NewUser(id int, firstName, lastName string) *User {
	return &User{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
	}
}

func (u User) String() string {
	return fmt.Sprintf("'[%v] %s %s'", u.Id, u.FirstName, u.LastName)
}
