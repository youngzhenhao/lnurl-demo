package api

func UploadUserInfo() {
	// TODO: Alice's upload workflow
	// 1. upload info to get LNURL

}

func PayToLnurl(ln string) {
	// TODO: pay-to-lnurl workflow by simulating Bob's operation
	// 1. decode LNURL
	// 2. send POST with amount by decoded URL to get invoice
	// 	  Use Same InvoiceResponse
	// 3. pay to invoice
	//	  config Bob's RPC_SERVER, TLS_CERT_PATH, MACAROON_PATH

}
