package db

import "testing"

func TestGetdb(t *testing.T) {
	db := getDB()
	err := db.Ping()
	if err != nil {
		t.Errorf("error connection to database %v", err.Error())
	}
	db.Close()
}