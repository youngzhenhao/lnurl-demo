package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"lnurl-demo/api"
	"lnurl-demo/boltDB"
	"time"
)

func main() {
	allInvoices()
	//boltDB.InitServerDB()
	//boltDB.InitPhoneDB()
}
func allInvoices() {
	db, err := bolt.Open("phone.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Printf("%s bolt.Open :%v\n", api.GetTimeNow(), err)
	}
	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("%s db.Close :%v\n", api.GetTimeNow(), err)
		}
	}(db)
	s := &boltDB.PhoneStore{DB: db}
	invoices, err := s.AllInvoices("invoices")
	if err != nil {
		return
	}
	fmt.Printf("%v\n", invoices)
}

func sampleOpenBoltDB() *boltDB.ServerStore {
	db, err := bolt.Open("store.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Printf("%s bolt.Open :%v\n", api.GetTimeNow(), err)
	}
	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("%s db.Close :%v\n", api.GetTimeNow(), err)
		}
	}(db)
	s := &boltDB.ServerStore{DB: db}
	return s
}
