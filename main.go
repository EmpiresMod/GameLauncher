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
	Test      string
)

func init() {

	flag.BoolVar(&Update, "u", true, "Check for updates")
	flag.StringVar(&UpdateURL, "url", "http://apollo.firebit.co.uk/~dc0/game", "The URL to fetch the manfiest and content from")
	flag.StringVar(&DirPath, "p", "./", "Path to the empires folder")
	flag.StringVar(&Content, "c", "", "The names of the content to download and install. Names should be seperated by a comma")
	flag.Parse()
}

func main() {

	if len(Content) != 0 {

		if Update {

			if err := UpdateManifest(); err != nil {

				log.Fatal(err)
			}
		}

		if err := ApplyAndLaunch(Content); err != nil {

			log.Fatal(err)
		}
	}

	if err := ShowGUI(); err != nil {

		log.Fatal(err)
	}
}

func UpdateManifest() (err error) {

	up := manifest.NewUpdate()
	up.Path = filepath.Join(DirPath, "manifest.json")
	up.URL = fmt.Sprintf("%s/%s/%s", UpdateURL, runtime.GOOS, "manifest.json")

	if err := up.Update(); err != nil {

		return err
	}

	return
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
