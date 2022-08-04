package common

import "os"

// IsFile returns true if given path is a file,
// or returns false when it's a directory or does not exist.
func IsFile(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
