package update

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	//"github.com/EmpiresMod/GameLauncher/signer"
	"github.com/EmpiresMod/GameLauncher/checksum"
)

const (
	FilePerm       = 0640
	DirectoryPerm  = 0750
	CheckSumPostix = ".sha256"
	TempPostix     = ".tmp"
)

// Error represents an Error reported in the HTTP session.
type Error struct {

	// Error message
	HTTPStatus string

	// The request sent to the server
	Request string
}

// Error returns a string representation of the HTTP error
func (e *Error) Error() string {

	return fmt.Sprintf("update: %s %s\n", e.HTTPStatus, e.Request)
}

type Update struct {

	// Path to file
	FileName string

	// URL of update
	URL string

	// Checksum of remote file
	CheckSum []byte
}

func New() *Update {

	return new(Update)
}

func (u *Update) Fetch() (b []byte, err error) {

	resp, err := http.Get(u.URL)
	if err != nil {

		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {

		return nil, &Error{resp.Status, u.URL}
	}

	return ioutil.ReadAll(resp.Body)
}

func (u *Update) GetCheckSum() ([]byte, error) {

	resp, err := http.Get(u.URL + CheckSumPostix)
	if err != nil {

		return nil, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {

		return nil, &Error{resp.Status, u.URL}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return nil, err
	}

	return bytes.Fields(b)[0], nil
}

// Check for updates to file. Returns true for updates and false for no updates
func (u *Update) Check() (ok bool, err error) {

	if !fileExists(u.FileName) {

		ok = true
		return
	}

	a, err := checksum.GenerateFileCheckSum(u.FileName)
	if err != nil {

		return
	}

	if len(u.CheckSum) == 0 {

		u.CheckSum, err = u.GetCheckSum()
		if err != nil {

			return
		}
	}

	if checksum.Compare(a, u.CheckSum) {

		ok = false
		return
	}

	ok = true
	return
}

func (u *Update) Update() (err error) {

	b, err := u.Fetch()
	if err != nil {

		return
	}

	if err = ioutil.WriteFile(u.FileName, b, FilePerm); err != nil {

		return
	}

	return
}
