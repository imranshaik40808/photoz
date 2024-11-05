package helper

import (
	"runtime"
	"strings"
)

func ChangePathSeparator(path string) string {
	os := runtime.GOOS
	if os == "windows" {
		return strings.ReplaceAll(path, "/", "\\")
	}
	path = strings.Replace(path, ":", "/", 1)
	return strings.ReplaceAll(path, "\\", "/")
}
