package lnurls

import (
	"encoding/json"
	"fmt"
	"io"
	"lnurl-demo/api"
	"net/http"
	"net/url"
)

type InvoiceResponse struct {
	Time    string `json:"time"`
	ID      string `json:"id"`
	Amount  int    `json:"amount"`
	Invoice string `json:"invoice"`
	Result  bool   `json:"result"`
}

func PostPhoneToAddInvoice(socket, amount string) string {
	targetUrl := "http://" + socket + "/addInvoice"

	payload := url.Values{"amount": {amount}}

	response, err := http.PostForm(targetUrl, payload)
	if err != nil {
		fmt.Printf("%s http.PostForm :%v\n", api.GetTimeNow(), err)
	}
	bodyBytes, _ := io.ReadAll(response.Body)

	var invoiceResponse InvoiceResponse
	if err := json.Unmarshal(bodyBytes, &invoiceResponse); err != nil {
		fmt.Printf("%s json.Unmarshal :%v\n", api.GetTimeNow(), err)
		return ""
	}
	return invoiceResponse.Invoice
}
