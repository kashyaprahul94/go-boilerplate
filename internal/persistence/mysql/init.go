package mysql

import (
	"fmt"
	"time"

	"github.com/kashyaprahul94/go-boilerplate/config"
	"github.com/kashyaprahul94/go-boilerplate/internal/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// mysqlInstance is the singleton database configured instance
var mysqlInstance *MySQLDatabaseInstance

// GetDatabaseInstance returns the MySQL database connection
func GetDatabaseInstance() *gorm.DB {
	return mysqlInstance.dbConnection
}

// initialize the database connection
func init() {
	appConfig := config.GetAppConfig()
	mysqlConfig := appConfig.Persistence.MySQL

	// Prepare the connection string
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DatabaseName)

	// Keep a single connection open and use it for all transactions
	dbConnection, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		logger.Error(fmt.Sprintf("Error while connecting to db :%s", err.Error()))
	}

	// Get the db instance from connection
	db, err := dbConnection.DB()

	if err != nil {
		defer db.Close()
		logger.Error(fmt.Sprintf("Error while connecting to db :%s", err.Error()))
	}

	db.SetConnMaxLifetime(10 * time.Second)

	if !appConfig.Environment.IsProduction() {
		// dbConnection.LogMode(true)
	}

	// Prepare the mysql instance
	mysqlInstance = &MySQLDatabaseInstance{
		host:         mysqlConfig.Host,
		port:         mysqlConfig.Port,
		user:         mysqlConfig.User,
		databaseName: mysqlConfig.DatabaseName,
		dbConnection: dbConnection,
	}
}
