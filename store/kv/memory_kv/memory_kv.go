package memory_kv

import (
	"errors"
	"fmt"
	"github.com/deffusion/chunkstore/store/kv"
	"github.com/syndtr/goleveldb/leveldb"
)

type MemoryKV struct {
	db map[string][]byte
}

func New() (kv.KV, error) {
	store := &MemoryKV{
		make(map[string][]byte),
	}
	return store, nil
}

func (m MemoryKV) closed() bool {
	return m.db == nil
}

func (m *MemoryKV) Close() (err error) {
	if m.closed() {
		err = errors.New(fmt.Sprint("MemoKV.Close: ", leveldb.ErrClosed))
		return
	}
	m.db = nil
	return nil
}

func (m *MemoryKV) Has(k []byte) (has bool, err error) {
	if m.closed() {
		err = errors.New(fmt.Sprintf("MemoKV.Has(%s): %s", k, leveldb.ErrClosed))
		return
	}
	_, has = m.db[string(k)]
	return
}

func (m *MemoryKV) Get(k []byte) (v []byte, err error) {
	errPrefix := fmt.Sprintf("MemoryKV.Get(%s): ", k)
	if m.closed() {
		err = errors.New(fmt.Sprint(errPrefix, leveldb.ErrClosed))
		return
	}
	v, ok := m.db[string(k)]
	if !ok {
		err = errors.New(fmt.Sprint(errPrefix, leveldb.ErrNotFound))
		return
	}
	return
}

func (m *MemoryKV) Put(k []byte, v []byte) (err error) {
	if m.closed() {
		err = errors.New(fmt.Sprintf("MemoryKV.Put(%s, %s): %s", k, v, leveldb.ErrClosed))
		return
	}
	m.db[string(k)] = v
	return
}
