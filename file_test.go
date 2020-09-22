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
}
