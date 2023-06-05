package merkle

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/deffusion/chunkstore/digest"
)

const Batch = 50

func Root(digests []digest.Digest) (digest.Digest, error) {
	errPrefix := "merkle.Root:"
	if len(digests) == 0 {
		return digest.Null, errors.New(fmt.Sprint(errPrefix, "hash an empty digest list"))
	}
	digestBytes := make([][]byte, 0, len(digests))
	prefix := digests[0].HashFuncPrefix()
	h, err := digests[0].HashFunc()
	if err != nil {
		return digest.Null, errors.New(fmt.Sprint(errPrefix, err))
	}
	for _, d := range digests {
		digestBytes = append(digestBytes, d.Bytes())
	}
	round := 0
	for len(digestBytes) > 1 {
		var next [][]byte
	Loop:
		for {
			h.Reset()
			for i := 0; i < Batch; i++ {
				if round*Batch+i == len(digestBytes) {
					next = append(next, h.Sum(nil))
					break Loop
				}
				h.Write(digestBytes[round*Batch+i])
			}
			next = append(next, h.Sum(nil))
			round++
		}
		digestBytes = next
	}
	d, err := digest.New(bytes.Join([][]byte{[]byte(prefix), digestBytes[0]}, nil))
	if err != nil {
		return digest.Null, errors.New(fmt.Sprint(errPrefix, err))
	}
	return d, nil
}

//func concat(bs [][]byte) []byte {
//	length := len(bs)
//	if length == 0 {
//		return nil
//	}
//	size := len(bs[0])
//	buff := bytes.NewBuffer(make([]byte, Batch*size, Batch*size))
//	for i, b := range bs {
//
//	}
//}
