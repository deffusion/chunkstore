package store

import (
	"fmt"
	"github.com/deffusion/chunkstore/digest"
	"github.com/deffusion/chunkstore/digest/digest_hash"
	"github.com/deffusion/chunkstore/merkle"
	"github.com/deffusion/chunkstore/splitter"
	"github.com/deffusion/chunkstore/store/kv"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	"os"
)

type ChunkStore struct {
	db        kv.KV
	chunkRoot string
	logger    *zap.Logger
}

func New(db kv.KV, chunkRoot string, logger *zap.Logger) *ChunkStore {
	return &ChunkStore{
		db,
		chunkRoot,
		logger.Named("ChunkStore"),
	}
}

func (cs *ChunkStore) Close() error {
	return cs.db.Close()
}

// Get returns digest of chunks of the given file digest
func (cs *ChunkStore) Get(d digest.Digest) ([]digest.Digest, error) {
	errorMessage := "ChunkStore.Get"
	data, err := cs.db.Get([]byte(d.String()))
	if err != nil {
		return nil, errors.WithMessage(err, errorMessage)
	}

	dl, err := digest.DecodeList(data)
	if err != nil {
		return nil, errors.WithMessage(err, errorMessage)
	}

	digests, err := dl.Digests()
	if err != nil {
		return nil, errors.WithMessage(err, errorMessage)
	}
	return digests, nil
}

// Add the file in the given path into chunkstore
func (cs *ChunkStore) Add(reader io.Reader) (digest.Digest, error) {
	errorMessage := "ChunkStore.Add"

	digests, err := splitter.SplitIntoFiles(ChunkRoot, reader, digest_hash.SHA256)
	if err != nil {
		return digest.Null, errors.WithMessage(err, errorMessage)
	}

	root, err := merkle.Root(digests)
	if err != nil {
		return digest.Null, errors.WithMessage(err, errorMessage)
	}

	dl := digest.ListFromDigests(digests)
	data, err := digest.EncodeList(dl)
	if err != nil {
		return digest.Null, errors.WithMessage(err, errorMessage)
	}

	err = cs.db.Put([]byte(root.String()), data)
	if err != nil {
		return digest.Null, errors.WithMessage(err, errorMessage)
	}

	return root, nil
}

// Extract write the file of the given digest to path
func (cs *ChunkStore) Extract(d digest.Digest, path string) error {
	logger := cs.logger.Named("Extract")
	errorMessage := "ChunkStore.Extract"

	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return errors.WithMessage(err, errorMessage)
	} else {
		logger.Info(fmt.Sprint("create file:", path))
	}

	digests, err := cs.Get(d)
	if err != nil {
		return errors.WithMessage(err, errorMessage)
	}

	for _, di := range digests {
		chunkFile, err := os.Open(fmt.Sprint(ChunkRoot, di.String()))
		if err != nil {
			chunkFile.Close()
			return errors.WithMessage(err, errorMessage)
		}

		n, err := io.Copy(file, chunkFile)
		chunkFile.Close()
		if err != nil {
			return errors.WithMessage(err, errorMessage)
		}

		if err != nil && err != io.EOF {
			return errors.WithMessage(err, fmt.Sprintf("%s (%d bytes were wrote)", errorMessage, n))
		}
	}
	return nil
}
