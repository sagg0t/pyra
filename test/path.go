package test

import (
	"os"
	"testing"
)

func SetCWDToProjectRoot(t *testing.T) {
	t.Chdir(os.Getenv("PROJECT_ROOT"))
}
