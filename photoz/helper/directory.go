package helper

import (
	"os"
	"path/filepath"
)

const ROOT = "photoz"

func CreateDirectory(folderPath string) bool {
	isExists, err := IsExists(folderPath)
	println(isExists)
	println(folderPath)
	if err != nil {
		panic(err)
	}
	if !isExists {
		err = os.Mkdir(folderPath, 0755)
		if err != nil {
			panic(err)
		}
		return true
	}
	return false
}

func CreateRootDirectory(folderPath string) bool {
	folderPath = filepath.Join(folderPath, ROOT)
	return CreateDirectory(folderPath)
}

func IsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
