package fileutil

import "os"

// Exist check if the file with the given path exits.
func Exist(path string) bool {
	fi, err := os.Lstat(path)
	if fi != nil || (err != nil && !os.IsNotExist(err)) {
		return true
	}

	return false
}
