package gofast

import (
	"testing"
)

func TestGetProjectRootPath(t *testing.T) {
	pwd, err := GetProjectRootPath()
	if err != nil {
		t.Error(err)
	}
	t.Logf("pwd: %v", pwd)
	pwd1, err := GetProjectRootGit()
	if pwd1 != pwd {
		t.Errorf("error GetProjectRootGit, got %v, want %v", pwd1, pwd)
	}
}
