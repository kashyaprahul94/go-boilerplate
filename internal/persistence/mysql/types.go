package mysql

import "gorm.io/gorm"

// MySQLDatabaseInstance represents the connected MySQL instance
type MySQLDatabaseInstance struct {
	// host of the connected database connection
	host string
	// port of the connected database connection
	port string

	// user of the connected database connection
	user string

	// database name of the connected database connection
	databaseName string

	// connected database connection
	dbConnection *gorm.DB
}
