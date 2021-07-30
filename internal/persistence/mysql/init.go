package mysql

import (
	"fmt"
	"go-boilerplate/config"
)

var db *MySQLDatabaseInstance

func DB() *MySQLDatabaseInstance {
	return db
}

func init() {
	config := config.GetAppConfig().DB.MySQL

	host := config.Host
	port := config.Port

	db = &MySQLDatabaseInstance{
		Host: host,
		Port: port,
		Conn: fmt.Sprintf("%s:%s", host, port),
	}

	fmt.Printf("Connected to DB: %v\n\n", db.Conn)
}
