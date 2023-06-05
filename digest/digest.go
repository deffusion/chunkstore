package digest

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/deffusion/chunkstore/digest/digest_hash"
	"hash"
)

type Digest struct {
	str string
}

var Null = Digest{}

func New(val interface{}) (d Digest, err error) {
	switch v := val.(type) {
	case string:
		d.str = v
	case []byte:
		d.str = fmt.Sprintf("%s%x", v[0:2], v[2:])
	}
	if _, err := d.HashFunc(); err != nil {
		return Null, errors.New(fmt.Sprintf("digest.New: %s", err))
	}
	_, err = hex.DecodeString(d.str[2:])
	if err != nil {
		return Null, errors.New(fmt.Sprintf("digest.New: digets decoding: %s", err))
	}
	return d, nil
}
func (d Digest) String() string {
	return d.str
}

// Bytes returns without the hash function prefix
func (d Digest) Bytes() (bytes []byte) {
	bytes, _ = hex.DecodeString(d.str[2:])
	return
}
func (d Digest) HashFunc() (hash.Hash, error) {
	return digest_hash.New(digest_hash.Hash(d.str))
}
func (d Digest) HashFuncPrefix() digest_hash.Hash {
	return digest_hash.Hash(d.str[0:2])
}
