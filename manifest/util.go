package manifest

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func getRemoteFileSize(url string) (int64, error) {

	resp, err := http.Get(url)
	if err != nil {

		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {

		return 0, fmt.Errorf("Remote server returned status code >= 400: %s", resp.Status)
	}

	return resp.ContentLength, nil
}

func getRemoteFile(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {

		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {

		return nil, errors.New(fmt.Sprintf("Remote server returned status code >= 400: %s", resp.Status))
	}

	return ioutil.ReadAll(resp.Body)
}

func getFileSize(path string) (int64, error) {

	f, err := os.Stat(path)
	if err != nil {

		return 0, err
	}

	return f.Size(), nil
}

func fileExists(path string) bool {

	if _, err := os.Stat(path); os.IsNotExist(err) {

		return false
	}

	return true
}

func fileHash(path string) (hash string, err error) {

	h := sha256.New()

	f, err := os.Open(path)
	if err != nil {

		return "", err
	}
	defer f.Close()

	if _, err = io.Copy(h, f); err != nil {

		return
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
