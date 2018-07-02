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

func File_not_exists_to_create(file string) error {
	if File_exists(file) {
		return nil
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	} else {
		f.Close()
	}

	return nil
}
