package mysql

import "fmt"

// Like an ORM API
func (db MySQLDatabaseInstance) InsertIntoTable(tableName string, params interface{}) {
	fmt.Printf("Should be able to insert %v into %s\n\n", params, tableName)
}
