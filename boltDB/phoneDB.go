package boltDB

func InitPhoneDB() error {
	_, err := createBucketInDB("phone.db", "invoices")
	return err
}
