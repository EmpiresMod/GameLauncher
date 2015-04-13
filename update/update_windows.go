package update

import (
	"io/ioutil"
	"os"
)

func (u *Update) UpdateExec() (err error) {

	b, err := u.Fetch()
	if err != nil {

		return
	}

	if fileExists(u.FileName + TempPostix) {

		if err = os.Remove(u.FileName + TempPostix); err != nil {

			return
		}
	}

	if err = os.Rename(u.FileName, u.FileName+TempPostix); err != nil {

		return
	}

	if err = ioutil.WriteFile(u.FileName, b, FilePerm); err != nil {

		return
	}

	return
}
