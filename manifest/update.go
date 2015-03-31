package manifest

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Update struct {

	// Path to file
	TargetPath string

	// URL of update
	TargetURL string

	// Hash of file
	Checksum string
}

func NewUpdate() *Update {

	return new(Update)
}

func (u *Update) Fetch() (b []byte, err error) {

	resp, err := http.Get(u.TargetURL)
	if err != nil {

		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {

		return nil, errors.New(fmt.Sprintf("Remote server returned status code >= 400: %s", resp.Status))
	}

	return ioutil.ReadAll(resp.Body)
}

func (u *Update) GetFileSize() (size int64, err error) {

	f, err := os.Stat(u.TargetPath)
	if err != nil {

		return
	}

	return f.Size(), nil
}

func (u *Update) GetRemoteSize() (int64, error) {

	resp, err := http.Get(u.TargetURL)
	if err != nil {

		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {

		return 0, fmt.Errorf("Remote server returned status code >= 400: %s", resp.Status)
	}

	return resp.ContentLength, nil
}

func (u *Update) Update() (err error) {

	if !fileExists(u.TargetPath) {

		b, err := u.Fetch()
		if err != nil {

			return err
		}

		if err = os.MkdirAll(filepath.Dir(u.TargetPath), DirectoryPerm); err != nil {

			return err
		}

		if err = ioutil.WriteFile(u.TargetPath, b, FilePerm); err != nil {

			return err
		}

		return nil
	}

	if len(u.Checksum) == 0 {

		size, err := u.GetFileSize()
		if err != nil {

			return err
		}

		rsize, err := u.GetRemoteSize()
		if err != nil {

			return err
		}

		if size != rsize {

			b, err := u.Fetch()
			if err != nil {

				return err
			}

			if err = ioutil.WriteFile(u.TargetPath, b, FilePerm); err != nil {

				return err
			}
		}

		return nil
	}

	hash, err := GenerateFileCheckSum(u.TargetPath)
	if err != nil {

		return err
	}

	if CompareCheckSums(hash, []byte(u.Checksum)) {

		b, err := u.Fetch()
		if err != nil {

			return err
		}

		if err = ioutil.WriteFile(u.TargetPath, b, FilePerm); err != nil {

			return err
		}
	}

	return
}
