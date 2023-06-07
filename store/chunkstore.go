package store

import (
	"fmt"
	"github.com/deffusion/chunkstore/digest"
	"github.com/deffusion/chunkstore/digest/digest_hash"
	"github.com/deffusion/chunkstore/merkle"
	"github.com/deffusion/chunkstore/splitter"
	"github.com/deffusion/chunkstore/store/kv"
	"go.uber.org/zap"
	"io"
	"os"
)

type ChunkStore struct {
	db        kv.KV
	chunkRoot string
	logger    *zap.Logger
}

func New(db kv.KV, chunkRoot string) *ChunkStore {
	logger, _ := zap.NewProduction()
	return &ChunkStore{
		db,
		chunkRoot,
		logger.Named("ChunkStore"),
	}
}

func (cs *ChunkStore) Close() error {
	return cs.db.Close()
}

// get
func (cs *ChunkStore) Get(d digest.Digest) []digest.Digest {
	errPrefix := "ChunkStore.Get: "
	data, err := cs.db.Get([]byte(d.String()))
	if err != nil {
		cs.logger.Error(fmt.Sprint(errPrefix, err))
		return nil
	}
	dl, err := digest.DecodeList(data)
	if err != nil {
		cs.logger.Error(fmt.Sprint(errPrefix, err))
		return nil
	}
	digests, err := dl.Digests()
	if err != nil {
		cs.logger.Error(fmt.Sprint(errPrefix, err))
		return nil
	}
	return digests
}

// Add the file in the given path into chunkstore
func (cs *ChunkStore) Add(file *os.File) digest.Digest {
	errPrefix := "ChunkStore.Add: "
	digests, err := splitter.SplitIntoFiles(ChunkRoot, file, digest_hash.SHA256)
	if err != nil {
		cs.logger.Error(fmt.Sprint(errPrefix, err))
		return digest.Null
	}
	root, err := merkle.Root(digests)
	if err != nil {
		cs.logger.Error(fmt.Sprint(errPrefix, err))
		return digest.Null
	}
	dl := digest.ListFromDigests(digests)
	data, err := digest.EncodeList(dl)
	if err != nil {
		cs.logger.Error(fmt.Sprint(errPrefix, err))
		return digest.Null
	}
	err = cs.db.Put([]byte(root.String()), data)
	if err != nil {
		cs.logger.Error(fmt.Sprint(errPrefix, err))
		return digest.Null
	}
	return root
}

func (cs *ChunkStore) Extract(d digest.Digest, path string) {
	logger := cs.logger.Named("Extract")
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		logger.Fatal(err.Error())
	} else {
		logger.Info(fmt.Sprint("create file:", path))
	}
	digests := cs.Get(d)
	for _, di := range digests {
		chunkFile, err := os.Open(fmt.Sprint(ChunkRoot, di.String()))
		n, err := io.Copy(file, chunkFile)
		chunkFile.Close()
		if err != nil && err != io.EOF {
			logger.Fatal(fmt.Sprintf("Chunkstore.Extract: %d bytes were wrote\n%s", n, err))
		}
	}
}
