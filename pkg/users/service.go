package users

// Service layer should expose low level APIs to access/modify domain model
func (u *User) InsertIntoMySql() {

	// Insert the user into mysql
	insertIntoMySql(u)
}


