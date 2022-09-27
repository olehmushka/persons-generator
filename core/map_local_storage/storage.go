package map_local_storage

import (
	"sync"
)

var m = sync.Map{}

func Save(key string, value []byte) {
	m.Store(key, value)
}

func Read(key string) []byte {
	v, ok := m.Load(key)
	if !ok {
		return nil
	}

	out, ok := v.([]byte)
	if !ok {
		return nil
	}

	return out
}

func Delete(key string) {
	m.Delete(key)
}
