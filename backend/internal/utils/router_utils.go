package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetProjectRootPath locates the project root directory by finding the go.mod file
// and appends any provided subdirectories to the path.
func GetProjectRootPath(subDirs ...string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %w", err)
	}

	rootDir, err := locateProjectRoot(cwd)
	if err != nil {
		return "", err
	}

	return filepath.Join(append([]string{rootDir}, subDirs...)...), nil
}

// locateProjectRoot traverses up the directory tree to find the project root by locating a go.mod file.
func locateProjectRoot(cwd string) (string, error) {
	for {
		if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err != nil {
			if os.IsNotExist(err) {
				parentDir := filepath.Dir(cwd)
				if parentDir == cwd {
					return "", fmt.Errorf("go.mod not found starting from directory: %s", cwd)
				}
				cwd = parentDir
			} else {
				return "", fmt.Errorf("error accessing path: %w", err)
			}
		} else {
			return cwd, nil
		}
	}
}
