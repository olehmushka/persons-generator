package json_storage

type Storage interface {
	Get(filename string) ([]byte, error)
	Store(filename string, file []byte) error
}
