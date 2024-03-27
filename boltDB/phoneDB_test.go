package boltDB

import "testing"

func TestInitPhoneDB(t *testing.T) {
	err := InitPhoneDB()
	if err != nil {
		return
	}
}
