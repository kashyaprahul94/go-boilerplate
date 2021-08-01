package users

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// GetUsers helps you to get the users from various data sources.
// For now, it returns all the users from MySQL database.
// In case of errors, it returns the User Domain Error.
func GetUsers() ([]*User, *UserDomainError) {
	mysqlUsers, err := getUsersFromMySql()

	if err != nil {
		return nil, UserMiscError(err.Error(), nil)
	}

	users := make([]*User, 0)

	// Map the MySQL user to domain User
	for _, mysqlUser := range mysqlUsers {
		users = append(users, &mysqlUser.User)
	}

	return users, nil
}

// GetUserById helps you to get the user by user id, from various data sources.
// For now, it returns user from MySQL database.
// In case of errors, it returns the User Domain Error.
func GetUserById(id int) (*User, *UserDomainError) {
	mysqlUser, err := getUserByIdFromMySql(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, UserNotFoundError(fmt.Sprintf("User with user id - %v, does not exist", id), nil)
		}

		return nil, UserMiscError(err.Error(), nil)
	}

	return &mysqlUser.User, nil
}

// Insert helps you to insert the user record to various data sources
// For now, it creates a user record inside MySQL database.
// In case of errors, it returns the User Domain Error.
func (u *User) Insert() (*User, *UserDomainError) {
	mysqlUser, err := insertUserIntoMySql(u)

	if err != nil {
		return nil, UserMiscError(err.Error(), nil)
	}

	return &mysqlUser.User, nil
}

// Upsert helps you to upsert the user record to various data sources
// For now, it creates a user record inside MySQL database.
// In case of errors, it returns the User Domain Error.
func (u *User) Upsert() (*User, *UserDomainError) {
	mysqlUser, err := upsertUserIntoMySql(u)

	if err != nil {
		return nil, UserMiscError(err.Error(), nil)
	}

	return &mysqlUser.User, nil
}

// PartialUpdate helps you to partially update non empty fields to the records
// For now, it partially updates a user record inside MySQL database.
// In case of errors, it returns the User Domain Error.
func (u *User) PartialUpdate() (*User, *UserDomainError) {
	mysqlUser, err := partialUpdateUserIntoMySql(u)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, UserNotFoundError(fmt.Sprintf("User with user id - %v, does not exist", u.ID), nil)
		}

		return nil, UserMiscError(err.Error(), nil)
	}

	return &mysqlUser.User, nil
}

// DeleteUserById helps you to delete the user by user id, from various data sources.
// For now, it deletes the user user from MySQL database.
// In case of errors, it returns the User Domain Error.
func DeleteUserById(id int) (*User, *UserDomainError) {
	mysqlUser, err := deleteUserByIdFromMySql(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, UserNotFoundError(fmt.Sprintf("User with user id - %v, does not exist", id), nil)
		}

		return nil, UserMiscError(err.Error(), nil)
	}

	return &mysqlUser.User, nil
}
