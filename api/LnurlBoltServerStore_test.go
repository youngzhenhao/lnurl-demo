package api

import (
	"fmt"
	"github.com/boltdb/bolt"
	"reflect"
	"testing"
	"time"
)

func TestInitServerDB(t *testing.T) {
	err := InitServerDB()
	if err != nil {
		return
	}
}

func TestBoltServerStore_AllUsers(t *testing.T) {
	db, err := bolt.Open("server.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Printf("%s bolt.Open :%v\n", GetTimeNow(), err)
	}

	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("%s db.Close :%v\n", GetTimeNow(), err)
		}
	}(db)
	s := &ServerStore{DB: db}
	users, err := s.AllUsers("users")
	if err != nil {
		return
	}
	got := len(users)
	want := 3
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestBoltServerStore_CreateUser(t *testing.T) {
	db, err := bolt.Open("server.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Printf("%s bolt.Open :%v\n", GetTimeNow(), err)
	}

	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("%s db.Close :%v\n", GetTimeNow(), err)
		}
	}(db)
	s := &ServerStore{DB: db}
	err = s.CreateOrUpdateUser("users", &User{
		ID:     "1",
		Name:   "1",
		Socket: "1",
	})
	if err != nil {
		return
	}

}

func TestBoltServerStore_UpdateUser(t *testing.T) {
	db, err := bolt.Open("server.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Printf("%s bolt.Open :%v\n", GetTimeNow(), err)
	}

	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("%s db.Close :%v\n", GetTimeNow(), err)
		}
	}(db)
	s := &ServerStore{DB: db}
	newUser := &User{
		ID:     "1",
		Name:   "111",
		Socket: "111",
	}
	err = s.CreateOrUpdateUser("users", newUser)
	if err != nil {
		return
	}
	user, err := s.ReadUser("users", "1")
	if err != nil {
		return
	}
	wantID := "1"
	if !reflect.DeepEqual(wantID, user.ID) {
		t.Errorf("expected:%v, got:%v", wantID, user.ID)
	}
	wantName := "111"
	if !reflect.DeepEqual(wantName, user.Name) {
		t.Errorf("expected:%v, got:%v", wantName, user.Name)
	}
	wantSocket := "111"
	if !reflect.DeepEqual(wantSocket, user.Socket) {
		t.Errorf("expected:%v, got:%v", wantSocket, user.Socket)
	}
	users, err := s.AllUsers("users")
	if err != nil {
		return
	}
	got := len(users)
	want := 4
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestBoltServerStore_ReadUser(t *testing.T) {
	db, err := bolt.Open("server.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Printf("%s bolt.Open :%v\n", GetTimeNow(), err)
	}

	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("%s db.Close :%v\n", GetTimeNow(), err)
		}
	}(db)
	s := &ServerStore{DB: db}
	user, err := s.ReadUser("users", "1")
	if err != nil {
		return
	}
	wantID := "1"
	if !reflect.DeepEqual(wantID, user.ID) {
		t.Errorf("expected:%v, got:%v", wantID, user.ID)
	}
	wantName := "1"
	if !reflect.DeepEqual(wantName, user.Name) {
		t.Errorf("expected:%v, got:%v", wantName, user.Name)
	}
	wantSocket := "1"
	if !reflect.DeepEqual(wantSocket, user.Socket) {
		t.Errorf("expected:%v, got:%v", wantSocket, user.Socket)
	}
}

func TestBoltServerStore_DeleteUser(t *testing.T) {
	db, err := bolt.Open("server.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Printf("%s bolt.Open :%v\n", GetTimeNow(), err)
	}

	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("%s db.Close :%v\n", GetTimeNow(), err)
		}
	}(db)
	s := &ServerStore{DB: db}

	err = s.DeleteUser("users", "3")
	if err != nil {
		return
	}
	user, _ := s.ReadUser("users", "3")
	if user != nil {
		return
	}
	users, err := s.AllUsers("users")
	if err != nil {
		return
	}
	got := len(users)
	want := 3
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}

}
