package main

import (
	"flag"
	"fmt"
	"github.com/boltdb/bolt"
	"lnurl-demo/api"
	"time"
)

func main() {

	//ListAllInvoices()
	//DecodeLnurl()
	PhoneRun()
	//UploadUserInfoRun()
	//PayToLnurlRun()
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

func PhoneRun() {
	api.RouterRunOnPhone()
}

func UploadUserInfoRun() {
	name := flag.String("name", "", "Alice's name")
	port := flag.String("port", "", "Alice's port")
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}
	fmt.Print(api.LnurlUploadUserInfo(*name, *port))
}

func PayToLnurlRun() {
	lnu := flag.String("lnu", "", "Alice's LNURL")
	amount := flag.String("amount", "", "Bob's amount pay to Alice")
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}
	fmt.Print(api.LnurlPayToLnu(*lnu, *amount))
}

func DecodeLnurl() {
	lnu := flag.String("lnu", "", "LNURL need to decode")
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}
	fmt.Print(api.Decode(*lnu))
}
