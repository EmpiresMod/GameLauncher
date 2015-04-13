package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/EmpiresMod/GameLauncher/manifest"
	"github.com/EmpiresMod/GameLauncher/update"
)

func UpdateManifest() (err error) {

	up := update.New()
	up.FileName = filepath.Join(DirPath, "manifest.json")
	up.URL = fmt.Sprintf("%s/%s/%s", UpdateURL, runtime.GOOS, up.FileName)

	ok, err := up.Check()
	if err != nil {

		return
	}

	if ok {

		if err = up.Update(); err != nil {

			return err
		}
	}

	return
}

func UpdateExecutable() (err error) {

	up := update.New()
	up.FileName = filepath.Join(DirPath, "launcher.exe")
	up.URL = fmt.Sprintf("%s/%s/%s", UpdateURL, runtime.GOOS, up.FileName)

	ok, err := up.Check()
	if err != nil {

		return
	}

	if ok {

		if err := up.UpdateExec(); err != nil {

			return err
		}
	}

	return
}

func ApplyAndLaunch(c string) (err error) {

	m, err := manifest.GetManifest(filepath.Join(DirPath, "manifest.json"), DirPath, UpdateURL)
	if err != nil {

		return
	}

	for _, v := range strings.Split(c, ",") {

		if err = m.Apply(v); err != nil {

			return
		}
	}

	if err = LaunchEmpires(); err != nil {

		return
	}

	return
}
