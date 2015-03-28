package manifest

import (
	"io/ioutil"
)

type Update struct {

	// Path to file
	Path string

	// URL of update
	URL string

	// Hash of file
	Sha256 string
}

func NewUpdate() *Update {

	return new(Update)
}

func (u *Update) Update() (err error) {

	if !fileExists(u.Path) {

		b, err := getRemoteFile(u.URL)
		if err != nil {

			return err
		}

		if err = ioutil.WriteFile(u.Path, b, FilePerm); err != nil {

			return err
		}

		return nil
	}

	if len(u.Sha256) == 0 {

		size, err := getFileSize(u.Path)
		if err != nil {

			return err
		}

		rsize, err := getRemoteFileSize(u.URL)
		if err != nil {

			return err
		}

		if size != rsize {

			b, err := getRemoteFile(u.URL)
			if err != nil {

				return err
			}

			if err = ioutil.WriteFile(u.Path, b, FilePerm); err != nil {

				return err
			}
		}

		return nil
	}

	h, err := fileHash(u.Path)
	if err != nil {

		return err
	}
	if h != u.Sha256 {

		b, err := getRemoteFile(u.URL)
		if err != nil {

			return err
		}

		if err = ioutil.WriteFile(u.Path, b, FilePerm); err != nil {

			return err
		}
	}

	return
}
