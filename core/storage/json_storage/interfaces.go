package json_storage

type Storage interface {
	Get(string) ([]byte, error)
	Store(string, []byte) error
	MkDir(string) error
	IsDirExists(string) bool
	GetDirInnerFilenames(dirname string) ([]string, error)
}
