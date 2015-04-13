package main

import (
	"flag"
	"log"
)

var (
	// Debug/Verbose switch
	Verbose bool

	// Update switch
	Update bool

	// Where can updates be found?
	UpdateURL string

	// Directory path where files are updated.
	DirPath string

	// What content to load/update.
	Content string
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

			if err := UpdateExecutable(); err != nil {

				log.Fatal(err)
			}
		}

		if err := ApplyAndLaunch(Content); err != nil {

			log.Fatal(err)
		}

		return
	}

	if err := ShowGUI(); err != nil {

		log.Fatal(err)
	}
}
