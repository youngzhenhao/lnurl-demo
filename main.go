package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"lnurl-demo/api"
	"lnurl-demo/boltDB"
	"lnurl-demo/lnurls"
	"time"
)

func main() {
	//allUsers()
	//allInvoices()

	//lnurls.RouterRunOnServer()
	//lnurls.RouterRunOnPhone()

	fmt.Printf(lnurls.Decode("LNURL1DP68GUP69UHNZV3H9CCZUVPWXYARJVPCXQHHQCTE8A5KG0FSXCMRVERZXANZ6DP5VCUJ6DPHVD3Z6C35XEJZ6CE4VF3RGCMPVSCRVVP55AU0KL"))
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
