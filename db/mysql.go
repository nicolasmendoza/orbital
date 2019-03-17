package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	driverName = "mysql"
	userName   = "root"
	pwd        = "password"
	database   = "curiosity"
)

var Pool *sql.DB
var connection = fmt.Sprintf("%s:%s@/%s", userName, pwd, database)

func getDB() *sql.DB {
	//if Pool == nil{
	return initDB()
	//}
	//return Pool
}

// The returned DB is safe for concurrent use by multiple goroutines
// and maintains its own pool of idle connections. Thus, the Open
// function should be called just once. It is rarely necessary to
// close a DB.
func initDB() *sql.DB {
	// root:password@/curiosity
	db, err := sql.Open(driverName, connection)
	if err != nil {
		panic(err.Error())
	}
	// set DB initialization to Pool
	return db
}
