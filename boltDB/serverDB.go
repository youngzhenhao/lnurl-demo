package boltDB

func InitServerDB() error {
	_, err := createBucketInDB("server.db", "users")
	return err
}
