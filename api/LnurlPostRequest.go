package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type InvoiceResponse struct {
	Time    string `json:"time"`
	ID      string `json:"id"`
	Amount  string `json:"amount"`
	Invoice string `json:"invoice"`
	Result  bool   `json:"result"`
}

type UserResponse struct {
	Time   string `json:"time"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Socket string `json:"socket"`
	Result bool   `json:"result"`
	Lnurl  string `json:"lnurl"`
}

// PostPhoneToAddInvoice called by server
func PostPhoneToAddInvoice(socket, amount string) string {
	targetUrl := "http://" + socket + "/addInvoice"
	//@dev: test
	targetUrl = "http://202.79.173.41:6000/addInvoice"
	payload := url.Values{"amount": {amount}}

	response, err := http.PostForm(targetUrl, payload)
	if err != nil {
		fmt.Printf("%s http.PostForm :%v\n", GetTimeNow(), err)
	}
	bodyBytes, _ := io.ReadAll(response.Body)

	var invoiceResponse InvoiceResponse
	if err := json.Unmarshal(bodyBytes, &invoiceResponse); err != nil {
		fmt.Printf("%s PPTAI json.Unmarshal :%v\n", GetTimeNow(), err)
		return ""
	}
	return invoiceResponse.Invoice
}
