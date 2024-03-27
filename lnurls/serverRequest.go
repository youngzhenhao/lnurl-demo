package lnurls

import (
	"encoding/json"
	"fmt"
	"io"
	"lnurl-demo/api"
	"net/http"
	"net/url"
	"reflect"
)

type APIResponse struct {
	Time    string `json:"time"`
	ID      string `json:"id"`
	Amount  int    `json:"amount"`
	Invoice string `json:"invoice"`
	Result  bool   `json:"result"`
}

func PostPhoneToAddInvoice(ip, amount string) {
	targetUrl := "http://" + ip + "/addInvoice"

	payload := url.Values{"amount": {amount}}

	response, err := http.PostForm(targetUrl, payload)
	if err != nil {
		fmt.Printf("%s http.PostForm :%v\n", api.GetTimeNow(), err)
	}
	bodyBytes, _ := io.ReadAll(response.Body)

	var apiResponse APIResponse
	if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
		fmt.Printf("%s json.Unmarshal :%v\n", api.GetTimeNow(), err)
		return
	}

	fmt.Printf("API Response: %+v\n", apiResponse)
	fmt.Println(reflect.TypeOf(apiResponse))
}
