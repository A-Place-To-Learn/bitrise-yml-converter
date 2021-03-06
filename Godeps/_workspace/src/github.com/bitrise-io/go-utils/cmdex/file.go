package cmdex

import (
	"os"
	"strings"

	"github.com/bitrise-io/go-utils/pathutil"
)

// CopyFile ...
func CopyFile(src, dst string) error {
	args := []string{src, dst}
	return RunCommand("rsync", args...)
}

// CopyDir ...
func CopyDir(src, dst string, isOnlyContent bool) error {
	if isOnlyContent && !strings.HasSuffix(src, "/") {
		src = src + "/"
	}
	args := []string{"-ar", src, dst}
	return RunCommand("rsync", args...)
}

// RemoveDir ...
func RemoveDir(dirPth string) error {
	if exist, err := pathutil.IsPathExists(dirPth); err != nil {
		return err
	} else if exist {
		if err := os.RemoveAll(dirPth); err != nil {
			return err
		}
	}
	return nil
}

// RemoveFile ...
func RemoveFile(pth string) error {
	if exist, err := pathutil.IsPathExists(pth); err != nil {
		return err
	} else if exist {
		if err := os.Remove(pth); err != nil {
			return err
		}
	}
	return nil
}
