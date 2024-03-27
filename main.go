package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"io"
	"lnurl-demo/api"
	"lnurl-demo/boltDB"
	"lnurl-demo/lnurls"
	"strings"
	"time"
)

func main() {
	//allUsers()
	//allInvoices()
	//boltDB.InitServerDB()
	//boltDB.InitPhoneDB()
	//fmt.Println(uuid.New().String())
	lnurls.RouterRunOnServer()
	//lnurls.RouterRunOnPhone()
	//fmt.Println(lnurls.Decode("LNURL1DP68GUP69UHNZV3H9CCZUVPWXYARJVPCXQHHQCTE8A5KG0FJXUENSDFJ8YCJ6WPS8PJJ6DPJV5EJ6WR9VCMZ6VENV4JN2DTYX5MNJVMPVCQU33"))
	//testReadAll()
}
func allInvoices() {
	_ = boltDB.InitPhoneDB()
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
	if len(invoices) == 0 {
		fmt.Printf("%v\n", invoices)
	} else {
		for _, v := range invoices {
			fmt.Printf("%v\n", v)
		}
	}

}

func allUsers() {
	_ = boltDB.InitServerDB()
	db, err := bolt.Open("server.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
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
	users, err := s.AllUsers("users")
	if err != nil {
		return
	}
	if len(users) == 0 {
		fmt.Printf("%v\n", users)
	} else {
		for _, v := range users {
			fmt.Printf("%v\n", v)
		}
	}
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

func testReadAll() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")
	b, err := io.ReadAll(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", b)
}
