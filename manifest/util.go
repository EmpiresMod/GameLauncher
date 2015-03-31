package manifest

import "os"

func fileExists(path string) bool {

	if _, err := os.Stat(path); os.IsNotExist(err) {

		return false
	}

	return true
}
