package gofast

import (
	"errors"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetProjectRootPath DEPRECATED, use GetProjectRootGit instead,
// finds an ancestor dir that have dir .git, returns its absolute path.
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

// GetProjectRootGit returns absolute path of the project root dir
// (base on git cmd, the computer needs to have git installed)
func GetProjectRootGit() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(stdout)), nil
}
