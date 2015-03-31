package checksum

import (
	"bytes"
	"crypto/sha256"
	"io"
	"os"
)

func GenerateCheckSum(r io.Reader) (hash []byte, err error) {

	hasher := sha256.New()

	_, err = io.Copy(hasher, r)
	if err != nil {

		return
	}

	return hasher.Sum(nil), nil
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

	return bytes.Equal(a, b)
}
