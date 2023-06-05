package digest_hash

import (
	"crypto/sha256"
	"errors"
	"hash"
	"hash/fnv"
)

type Hash string

const (
	FNV32  Hash = "f5"
	FNV64  Hash = "f6"
	FNV128 Hash = "f7"
	SHA256 Hash = "s8"
)

func New(s Hash) (h hash.Hash, err error) {
	switch s[0:2] {
	case FNV32:
		h = fnv.New32()
	case FNV64:
		h = fnv.New64()
	case FNV128:
		h = fnv.New128()
	case SHA256:
		h = sha256.New()
	}
	if h == nil {
		err = errors.New("digest_hash.New: unsupported digest hash function type")
	}
	return
}
