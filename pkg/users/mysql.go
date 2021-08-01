package users

import (
	"github.com/kashyaprahul94/go-boilerplate/internal/persistence/mysql"

	"gorm.io/gorm/clause"
)

// A helper method for GORM specifying the table name to be used for User model.
func (MySqlUser) TableName() string {
	return "users"
}

// Get all users from MySQL user table
func getUsersFromMySql() ([]*MySqlUser, error) {
	db := mysql.GetDatabaseInstance()

	var users []*MySqlUser

	if result := db.Model(&MySqlUser{}).Find(&users); result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

// Get one user by id from MySQL user table
func getUserByIdFromMySql(userId int) (*MySqlUser, error) {
	db := mysql.GetDatabaseInstance()

	var user MySqlUser

	if result := db.Model(&MySqlUser{}).Where("id = ?", userId).First(&user); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// Insert one user into MySQL user table
func insertUserIntoMySql(user *User) (*MySqlUser, error) {
	db := mysql.GetDatabaseInstance()

	mysqlUser := &MySqlUser{
		User: *user,
	}

	if result := db.Model(&MySqlUser{}).Select("first_name", "last_name", "email").Create(&mysqlUser); result.Error != nil {
		return nil, result.Error
	}

	return mysqlUser, nil
}

// Upsert one user into MySQL user table
func upsertUserIntoMySql(user *User) (*MySqlUser, error) {
	db := mysql.GetDatabaseInstance()

	mysqlUser := &MySqlUser{
		User: *user,
	}

	if result := db.Model(&MySqlUser{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"first_name", "last_name", "email"}),
	}).Create(&mysqlUser); result.Error != nil {
		return nil, result.Error
	}

	return mysqlUser, nil
}

// Upsert one user into MySQL user table
func partialUpdateUserIntoMySql(user *User) (*MySqlUser, error) {
	db := mysql.GetDatabaseInstance()

	mysqlUser := &MySqlUser{
		User: *user,
	}

	if result := db.Model(&MySqlUser{}).Where("id = ?", mysqlUser.User.ID).Updates(&mysqlUser).First(&mysqlUser); result.Error != nil {
		return nil, result.Error
	}

	return mysqlUser, nil
}

// Delete one user by id from MySQL user table
func deleteUserByIdFromMySql(userId int) (*MySqlUser, error) {
	db := mysql.GetDatabaseInstance()

	mysqlUser := &MySqlUser{
		User: User{
			ID: uint(userId),
		},
	}

	if result := db.Model(&MySqlUser{}).Where("id = ?", userId).First(&mysqlUser).Delete(&mysqlUser); result.Error != nil {
		return nil, result.Error
	}

	return mysqlUser, nil
}
