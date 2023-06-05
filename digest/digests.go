package digest

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/deffusion/chunkstore/digest/digest_hash"
)

type DigestList struct {
	HashType    digest_hash.Hash
	DigestBytes [][]byte
}

func ListFromDigests(digests []Digest) *DigestList {
	l := &DigestList{}
	if len(digests) == 0 {
		return l
	}
	l.HashType = digests[0].HashFuncPrefix()
	l.DigestBytes = make([][]byte, len(digests))
	for i, d := range digests {
		l.DigestBytes[i] = d.Bytes()
	}
	return l
}

func (dl *DigestList) Digests() ([]Digest, error) {
	var err error
	ds := make([]Digest, len(dl.DigestBytes))
	for i, db := range dl.DigestBytes {
		ds[i], err = New(fmt.Sprintf("%s%x", dl.HashType, db))
		if err != nil {
			return nil, errors.New(fmt.Sprint("DigestList.Digests: ", err))
		}
	}
	return ds, nil
}

func EncodeList(dl *DigestList) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(*dl)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("EncodeList: %s", err))
	}
	return buf.Bytes(), nil
}

func DecodeList(data []byte) (*DigestList, error) {
	var dl DigestList
	var buf bytes.Buffer
	buf.Write(data)
	dec := gob.NewDecoder(&buf)
	err := dec.Decode(&dl)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("DecodeList: %s", err))
	}
	return &dl, nil
}
