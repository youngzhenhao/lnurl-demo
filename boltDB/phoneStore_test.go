package boltDB

import (
	"fmt"
	"github.com/boltdb/bolt"
	"lnurl-demo/api"
	"reflect"
	"testing"
	"time"
)

func TestInitPhoneDB(t *testing.T) {
	err := InitPhoneDB()
	if err != nil {
		return
	}
}

func TestPhoneStore_AllInvoices(t *testing.T) {
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
	s := &PhoneStore{DB: db}
	invoices, err := s.AllInvoices("invoices")
	if err != nil {
		return
	}
	got := len(invoices)
	want := 4
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestPhoneStore_CreateInvoice(t *testing.T) {
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
	s := &PhoneStore{DB: db}
	err = s.CreateOrUpdateInvoice("invoices", &Invoice{
		ID: "4",
		//PubKey:     "4",
		Amount:     44,
		InvoiceStr: "4",
	})
	if err != nil {
		return
	}
}

func TestPhoneStore_UpdateInvoice(t *testing.T) {
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
	s := &PhoneStore{DB: db}
	newInvoice := &Invoice{
		ID: "1",
		//PubKey:     "111",
		Amount:     111,
		InvoiceStr: "111",
	}
	err = s.CreateOrUpdateInvoice("invoices", newInvoice)
	if err != nil {
		return
	}
	invoice, err := s.ReadInvoice("invoices", "1")
	if err != nil {
		return
	}
	wantID := "1"
	if !reflect.DeepEqual(wantID, invoice.ID) {
		t.Errorf("expected:%v, got:%v", wantID, invoice.ID)
	}
	//wantPubKey := "111"
	//if !reflect.DeepEqual(wantPubKey, invoice.PubKey) {
	//	t.Errorf("expected:%v, got:%v", wantPubKey, invoice.PubKey)
	//}
	wantAmount := 111
	if !reflect.DeepEqual(wantAmount, invoice.Amount) {
		t.Errorf("expected:%v, got:%v", wantAmount, invoice.Amount)
	}
	wantInvoiceStr := "111"
	if !reflect.DeepEqual(wantInvoiceStr, invoice.InvoiceStr) {
		t.Errorf("expected:%v, got:%v", wantInvoiceStr, invoice.InvoiceStr)
	}
	invoices, err := s.AllInvoices("invoices")
	if err != nil {
		return
	}
	got := len(invoices)
	want := 4
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestPhoneStore_ReadInvoice(t *testing.T) {
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
	s := &PhoneStore{DB: db}
	invoice, err := s.ReadInvoice("invoices", "4")
	if err != nil {
		return
	}
	wantID := "4"
	if !reflect.DeepEqual(wantID, invoice.ID) {
		t.Errorf("expected:%v, got:%v", wantID, invoice.ID)
	}
	//wantPubKey := "4"
	//if !reflect.DeepEqual(wantPubKey, invoice.PubKey) {
	//	t.Errorf("expected:%v, got:%v", wantPubKey, invoice.PubKey)
	//}
	wantAmount := 44
	if !reflect.DeepEqual(wantAmount, invoice.Amount) {
		t.Errorf("expected:%v, got:%v", wantAmount, invoice.Amount)
	}
	wantInvoiceStr := "4"
	if !reflect.DeepEqual(wantInvoiceStr, invoice.InvoiceStr) {
		t.Errorf("expected:%v, got:%v", wantInvoiceStr, invoice.InvoiceStr)
	}
}

func TestPhoneStore_DeleteInvoice(t *testing.T) {
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
	s := &PhoneStore{DB: db}

	err = s.DeleteInvoice("invoices", "3")
	if err != nil {
		return
	}
	invoice, _ := s.ReadInvoice("invoices", "3")
	if invoice != nil {
		return
	}
	invoices, err := s.AllInvoices("invoices")
	if err != nil {
		return
	}
	got := len(invoices)
	want := 3
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}

}
