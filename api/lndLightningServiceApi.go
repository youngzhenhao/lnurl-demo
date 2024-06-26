package api

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/grpc"
)

// AddInvoiceImportEnv called by Alice
//
//	@Description:AddInvoice attempts to add a new invoice to the invoice database.
//	Any duplicated invoices are rejected, therefore all invoices must have a unique payment preimage.
//	@return string
func AddInvoiceImportEnv(value int64, memo, _rpcServer, _tlsCertPath, _macaroonPath string) string {
	grpcHost := GetEnv(_rpcServer)
	tlsCertPath := GetEnv(_tlsCertPath)
	macaroonPath := GetEnv(_macaroonPath)
	creds := newTlsCert(tlsCertPath)
	macaroon := getMacaroon(macaroonPath)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close err: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.Invoice{
		Value: value,
		Memo:  memo,
	}
	response, err := client.AddInvoice(context.Background(), request)
	if err != nil {
		fmt.Printf("%s client.AddInvoice :%v\n", GetTimeNow(), err)
		return ""
	}
	return response.GetPaymentRequest()
}

// SendPaymentSyncImportEnv called by Bob
//
//	@Description:SendPaymentSync is the synchronous non-streaming version of SendPayment.
//	This RPC is intended to be consumed by clients of the REST proxy. Additionally, this RPC expects the destination's public key and the payment hash (if any) to be encoded as hex strings.
//	@return string
func SendPaymentSyncImportEnv(invoice, _rpcServer, _tlsCertPath, _macaroonPath string) string {
	grpcHost := GetEnv(_rpcServer)
	tlsCertPath := GetEnv(_tlsCertPath)
	macaroonPath := GetEnv(_macaroonPath)
	creds := newTlsCert(tlsCertPath)
	macaroon := getMacaroon(macaroonPath)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close err: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.SendRequest{
		PaymentRequest: invoice,
	}
	stream, err := client.SendPaymentSync(context.Background(), request)
	if err != nil {
		fmt.Printf("%s client.SendPaymentSync :%v\n", GetTimeNow(), err)
		return "false"
	}
	return hex.EncodeToString(stream.PaymentHash)
}
