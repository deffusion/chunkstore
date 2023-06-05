package kv

type KV interface {
	Close() error
	Has([]byte) (bool, error)
	Get(k []byte) ([]byte, error)
	Put([]byte, []byte) error
}
