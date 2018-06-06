package tool

import (
	"os"
)

func File_exists(file string) bool {
	_, err := os.Stat(file)
	if err == nil {
		return true
	}
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			return true
		}
	}
	return false
}
