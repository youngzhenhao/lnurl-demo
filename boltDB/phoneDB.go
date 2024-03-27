package boltDB

func InitPhoneDB() error {
	_, err := createBucketInPhoneDB("./phone.db", "invoices")
	return err
}
