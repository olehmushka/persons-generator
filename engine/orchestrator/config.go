package orchestrator

type Config struct {
	StorageFolderName string
	RedisURL          string
	RedisUsername     string
	RedisPassword     string

	MongoDBURL              string
	MongoDBDBName           string
	MongoDBUsername         string
	MongoDBPassword         string
	MongoDBMaxBulkItemsSize int
}
