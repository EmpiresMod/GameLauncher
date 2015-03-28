package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/EmpiresMod/GameLauncher/manifest"
)

var (
	Update    bool
	UpdateURL string
	DirPath   string
	Content   string
)

func init() {

	flag.BoolVar(&Update, "u", true, "Check for updates.")
	flag.StringVar(&UpdateURL, "url", "https://apollo.firebit.co.uk/~dc0/game", "The URL to fetch updates from.")
	flag.StringVar(&DirPath, "p", "./", "Path to empires folder.")
	flag.StringVar(&Content, "e", "", "Specify the content to download and enable.")
	flag.Parse()
}

func main() {

	if Update {

		// Update the launcher
		/*up := manifest.NewUpdate()
		up.Path = filepath.Join(DirPath, "launcher.exe")
		up.URL = fmt.Sprintf("%s/%s/%s", UpdateURL, runtime.GOOS, "launcher.exe")

		if err := up.Update(); err != nil {

			log.Fatal(err)
		}*/

		// Update the manifest
		up := manifest.NewUpdate()
		up.Path = filepath.Join(DirPath, "manifest.json")
		up.URL = fmt.Sprintf("%s/%s/%s", UpdateURL, runtime.GOOS, "manifest.json")

		if err := up.Update(); err != nil {

			log.Fatal(err)
		}
	}

	if len(Content) != 0 {

		if err := ApplyAndLaunch(Content); err != nil {

			log.Fatal(err)
		}
	}

	if err := ShowGUI(); err != nil {

		log.Fatal(err)
	}
}

func ApplyAndLaunch(c string) (err error) {

	m, err := manifest.GetManifest(filepath.Join(DirPath, "manifest.json"), DirPath, UpdateURL)
	if err != nil {

		return
	}

	for _, v := range strings.Split(c+",Gameinfo", ",") {

		if err = m.Apply(v); err != nil {

			return
		}
	}

	if err = LaunchEmpires(); err != nil {

		return
	}

	return
}
