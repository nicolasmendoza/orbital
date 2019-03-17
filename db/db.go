package db

import "database/sql"

// By default connected to mysql driver.
func Get()*sql.DB{
	return getDB()
}
