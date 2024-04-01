package main

import (
	"flag"
	"fmt"
	"github.com/boltdb/bolt"
	"lnurl-demo/api"
	"time"
)

func main() {

	//ListAllUsers()
	//DecodeLnurl()

	ServerRun()

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

func ServerRun() {
	api.RouterRunOnServer()
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
