package cmd

import (
	"os"
	"path/filepath"
	"strings"
)

func expandTilde(path string) (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(path, "~/") {
		return filepath.Join(dir, (path)[1:]), nil
	}
	return path, nil
}
