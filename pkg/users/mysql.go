package users

import (
	"fmt"

	"github.com/kashyaprahul94/go-boilerplate/internal/persistence/mysql"
)

func insertIntoMySql(user *User) {
	mysqlUser := &User{
		Id: user.Id,
		FirstName: user.FirstName,
		LastName: user.LastName,
	}

	// Get DB Instance
	conn := mysql.DB()

	// Dummy insertion
	conn.InsertIntoTable("users", mysqlUser)


	fmt.Printf("User is saved to MySQL: %v\n\n", mysqlUser)
}
