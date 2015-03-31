package signer

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
	"io/ioutil"
	"os"

	"github.com/EmpiresMod/GameLauncher/checksum"
)

func ParsePublicPEM(b []byte) (pub *rsa.PublicKey, err error) {

	block, _ := pem.Decode(b)
	if block == nil {

		return nil, errors.New("Could not parse PEM data")
	}

	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {

		return
	}

	return key.(*rsa.PublicKey), nil
}

func ParseFilePublicPEM(filename string) (pub *rsa.PublicKey, err error) {

	b, err := ioutil.ReadFile(filename)
	if err != nil {

		return
	}

	return ParsePublicPEM(b)
}

func VerifyCryptoSignature(r io.Reader, sig []byte, pub *rsa.PublicKey) (err error) {

	hash, err := checksum.GenerateCheckSum(r)
	if err != nil {

		return
	}

	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hash, sig)
}

func VerifyFileCryptoSignature(filename string, sig []byte, pub *rsa.PublicKey) (err error) {

	f, err := os.Open(filename)
	if err != nil {

		return
	}
	defer f.Close()

	return VerifyCryptoSignature(f, sig, pub)
}
