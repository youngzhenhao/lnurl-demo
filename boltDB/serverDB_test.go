package boltDB

import "testing"

func TestInitServerDB(t *testing.T) {
	err := InitServerDB()
	if err != nil {
		return
	}
}
