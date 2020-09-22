package gofast

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// GetProjectRootPath finds an ancestor dir that have dir .git,
// this func returns an absolute path.
func GetProjectRootPath() (string, error) {
	const pivot = ".git"
	currentDir := "./"
	isFound := false
OuterFor:
	for i := 0; i < 32; i++ {
		files, err := ioutil.ReadDir(currentDir)
		if err != nil {
			return "", err
		}
		for _, file := range files {
			if file.Name() == pivot {
				isFound = true
				break OuterFor
			}
		}
		currentDir = "../" + currentDir
	}
	if !isFound {
		return "", errors.New(`max retries to find ".git"`)
	}
	absPath, err := filepath.Abs(currentDir)
	return absPath, err
}
