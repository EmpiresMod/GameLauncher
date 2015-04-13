package manifest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/EmpiresMod/GameLauncher/update"
)

const (
	DisabledPostfix = ".disabled"
	FilePerm        = 0640
	DirectoryPerm   = 0750
)

type Manifest struct {

	// Base URL
	BaseURL string `json:"-"`

	// Base file path
	BasePath string `json:"-"`

	// Encoded public key
	PemBytes []byte `json:"-"`

	// array holding defentions of files
	Files []File
}

type File struct {

	// name of custom content
	Name string

	// Is this a directory?
	IsDir bool

	// Can this file be disabled?
	CanDisable bool

	// array of other files this depends on
	Depends []string

	// array of conflicting custom content if any
	Conflicts []string

	// relative to the base path
	Path string

	// File hash sum
	CheckSum string
}

func (m *Manifest) Disable(name string) (err error) {

	r, err := regexp.Compile(name)
	if err != nil {

		return err
	}

	for _, v := range m.Files {

		if !v.CanDisable {

			continue
		}

		if !r.MatchString(v.Name) {

			continue
		}

		oldPath := filepath.Join(m.BasePath, v.Path)
		newPath := filepath.Join(m.BasePath, v.Path+DisabledPostfix)

		if !fileExists(oldPath) {

			continue
		}

		if err = os.Rename(oldPath, newPath); err != nil {

			return
		}
	}

	return
}

func (m *Manifest) Enable(name string) (err error) {

	r, err := regexp.Compile(name)
	if err != nil {

		return err
	}

	for _, v := range m.Files {

		if !v.CanDisable {

			continue
		}

		if !r.MatchString(v.Name) {

			continue
		}

		oldPath := filepath.Join(m.BasePath, v.Path+DisabledPostfix)
		newPath := filepath.Join(m.BasePath, v.Path)

		if !fileExists(oldPath) {

			continue
		}

		if err = os.Rename(oldPath, newPath); err != nil {

			return
		}
	}

	return
}

func (m *Manifest) Apply(name string) (err error) {

	for _, v := range m.Files {

		if v.Name != name {

			continue
		}

		// Is a directory?
		if v.IsDir {

			path := filepath.Join(m.BasePath, v.Path)
			if !fileExists(path) {

				if err = os.MkdirAll(path, DirectoryPerm); err != nil {

					return
				}
			}

			continue
		}

		// Resolve conflicts
		for _, vv := range v.Conflicts {

			if err = m.Disable(vv); err != nil {

				return
			}
		}

		// Check if its disabled
		if fileExists(filepath.Join(m.BasePath, v.Path+DisabledPostfix)) {

			if err = m.Enable(name); err != nil {

				return
			}
		}

		if len(v.Path) != 0 {

			up := update.New()
			up.FileName = filepath.Join(m.BasePath, v.Path)
			up.URL = fmt.Sprintf("%s/%s/%s", m.BaseURL, runtime.GOOS, v.Path)
			up.CheckSum = []byte(v.CheckSum)

			ok, err := up.Check()
			if err != nil {

				return err
			}

			if ok {

				if err = up.Update(); err != nil {

					return err
				}
			}
		}
	}

	return
}

func GetManifest(filename, path, url string) (m *Manifest, err error) {

	b, err := ioutil.ReadFile(filename)
	if err != nil {

		return
	}

	err = json.Unmarshal(b, &m)
	if err != nil {

		return
	}
	m.BasePath = path
	m.BaseURL = url

	return
}
