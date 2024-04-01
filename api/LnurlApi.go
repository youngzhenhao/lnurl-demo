package api

// LnurlUploadUserInfo
//
//	@Description: Alice's upload workflow, call after Alice and server's web services are launched
//	Alice's front-end uses this LNURL to generate a QR code that waits to be scanned
//	@param name
//	@param port
//	@return string
func LnurlUploadUserInfo(name, port string) string {
	// TODO: Alice's upload workflow
	// 1. upload info to get LNURL
	return PostServerToUploadUserInfo(name, port)
}

// LnurlPayToLnu
//
//	@Description: Bob's pay-to-lnurl workflow, call after Alice's LNURL QR code is generated
//	Bob's front-end scans the Alice's QR code to get the LNURL and then calls the LnurlPayToLnu with amount which Bob wanna pay
//	@param ln
//	@param amount
//	@return string
func LnurlPayToLnu(lnu, amount string) string {
	// TODO: pay-to-lnurl workflow by simulating Bob's operation
	// 0. decode LNURL(NEED NO MORE)
	// 1. send POST with amount by decoded URL to get invoice
	// 	  Use Same InvoiceResponse
	invoice := PostServerToPayByPhoneAddInvoice(lnu, amount)
	// 2. pay to invoice
	//	  config Bob's RPC_SERVER, TLS_CERT_PATH, MACAROON_PATH
	return SendPaymentSyncImportEnv(invoice, "BOB_RPC_SERVER", "BOB_TLS_CERT_PATH", "BOB_MACAROON_PATH")
}
