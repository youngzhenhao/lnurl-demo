package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"lnurl-demo/api"
	"lnurl-demo/boltDB"
	"time"
)

func main() {
	//allUsers()
	boltDB.InitServerDB()
	boltDB.InitPhoneDB()
}

func read() {
	//fmt.Printf(api.AddInvoice(100, ""))
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
	s := &boltDB.Store{DB: db}
	user, err := s.ReadUser("users", "1")
	if err != nil {
		return
	}
	fmt.Printf("%v\n", user)
}

func allUsers() {
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
	s := &boltDB.Store{DB: db}
	users, err := s.AllUsers("users")
	if err != nil {
		return
	}
	fmt.Printf("%v\n", users)
}

func sampleOpenBoltDB() *boltDB.Store {
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
	s := &boltDB.Store{DB: db}
	return s
}
