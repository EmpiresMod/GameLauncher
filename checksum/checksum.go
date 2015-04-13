package checksum

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func GenerateCheckSum(r io.Reader) (hash []byte, err error) {

	hasher := sha256.New()

	_, err = io.Copy(hasher, r)
	if err != nil {

		return
	}

	hash = make([]byte, sha256.BlockSize)
	hex.Encode(hash, hasher.Sum(nil))

	return hash, nil
}

func GenerateFileCheckSum(filename string) (hash []byte, err error) {

	f, err := os.Open(filename)
	if err != nil {

		return
	}
	defer f.Close()

	return GenerateCheckSum(f)
}

func Compare(a, b []byte) (ok bool) {

	return bytes.EqualFold(a, b)
}
