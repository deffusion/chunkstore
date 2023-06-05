package level_kv

import (
	"errors"
	"fmt"
	"github.com/deffusion/chunkstore/store/kv"
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelKV struct {
	path string
	db   *leveldb.DB
}

func New(path string) (kv.KV, error) {
	var err error
	store := &LevelKV{
		path: path,
	}

	store.db, err = leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprint("LevelKV.New:", err))
	}

	return store, nil
}

func (ls *LevelKV) Close() error {
	return ls.db.Close()
}

func (ls *LevelKV) Has(k []byte) (ret bool, err error) {
	ret, err = ls.db.Has(k, nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("LevelKV.Has(%s): %s", k, err))
	}
	return
}

func (ls *LevelKV) Get(k []byte) (val []byte, err error) {
	val, err = ls.db.Get(k, nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("LevelKV.Get(%s): %s", k, err))
	}
	return
}

func (ls *LevelKV) Put(k, v []byte) (err error) {
	err = ls.db.Put(k, v, nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("LevelKV.Put(%s, %s): %s", k, v, err))
	}
	return
}
