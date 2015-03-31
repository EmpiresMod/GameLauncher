package manifest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/EmpiresMod/GameLauncher/checksum"
	//"github.com/EmpiresMod/GameLauncher/signer"
	"github.com/EmpiresMod/GameLauncher/update"
)

const (
	DisabledPostfix = ".disabled"
	FilePerm        = 0640
	DirectoryPerm   = 0750
)

type Manifest struct {

	// Base URL
	baseURL string

	// Base file path
	basePath string

	// Encoded public key
	pemBytes []byte

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
	Checksum string
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

		oldPath := filepath.Join(m.basePath, v.Path)
		newPath := filepath.Join(m.basePath, v.Path+DisabledPostfix)

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

		oldPath := filepath.Join(m.basePath, v.Path+DisabledPostfix)
		newPath := filepath.Join(m.basePath, v.Path)

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

			path := filepath.Join(m.basePath, v.Path)
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
		if fileExists(filepath.Join(m.basePath, v.Path+DisabledPostfix)) {

			if err = m.Enable(name); err != nil {

				return
			}
		}

		if len(v.Path) != 0 {

			if fileExists(filepath.Join(m.basePath, v.Path)) {

				hash, err := checksum.GenerateFileCheckSum(v.Path)
				if err != nil {

					return err
				}

				if checksum.Compare([]byte(v.Checksum), hash) {

					continue
				}

			}

			up := update.New()
			up.TargetPath = filepath.Join(m.basePath, v.Path)
			up.TargetURL = fmt.Sprintf("%s/%s/%s", m.baseURL, runtime.GOOS, v.Path)

			if err = up.Update(); err != nil {

				return
			}
		}
	}

	return
}

func GetManifest(path, basepath, baseurl string) (m *Manifest, err error) {

	b, err := ioutil.ReadFile(path)
	if err != nil {

		return
	}

	err = json.Unmarshal(b, &m)
	if err != nil {

		return
	}
	m.basePath = basepath
	m.baseURL = baseurl

	return
}
