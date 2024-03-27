package boltDB

func InitServerDB() error {
	_, err := createBucketInServerDB("./server.db", "users")
	return err
}
