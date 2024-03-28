package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"lnurl-demo/api"
	"time"
)

func main() {

	//api.RouterRunOnServer()
	//api.RouterRunOnPhone()

	// TODO: Multiple-Command CLI
	//ListAllUsers()
	//ListAllInvoices()

	//lnu := flag.String("lnu", "", "lnurl need to decode")
	//flag.Parse()
	//if flag.NFlag() == 0 {
	//	flag.Usage()
	//	return
	//}
	//fmt.Println(api.Decode(*lnu))
}

func ListAllInvoices() {
	_ = api.InitPhoneDB()
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
	s := &api.PhoneStore{DB: db}
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

func ListAllUsers() {
	_ = api.InitServerDB()
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
	s := &api.ServerStore{DB: db}
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
