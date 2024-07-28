package main

import (
	"os"
	"path/filepath"
)

func ProcessAndCleanupTempFile(tempFile *os.File) {
	// Logic processing
	// Clean up
	defer os.Remove(tempFile.Name())
}
func ProcessAndCleanupTempDir(tempDir string) {
	// Logic processing
	// Clean up
	defer os.RemoveAll(tempDir)
}

func CreateTempFile(prefix string) (*os.File, error) {
	tempFile, err := os.CreateTemp("", prefix)
	if err != nil {
		return nil, err
	}
	// CreateTemp creates file with os.O_RDWR|os.O_CREATE|os.O_EXCL mode
	return tempFile, nil
}

func CreateTempDir(prefix string) (string, error) {
	tempDir, err := os.MkdirTemp("", prefix)
	if err != nil {
		return "", err
	}
	return tempDir, nil
}

func CleanUpEmptyDirectories(rootDir string) error {
	return filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			entries, err := os.ReadDir(path)
			if err != nil {
				return err
			}
			if len(entries) == 0 && path != rootDir {
				if err := os.Remove(path); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func ArrangeBooksByAuthor(libPath string, books []Book) error {
	for _, book := range books {
		authorDir := filepath.Join(libPath, SanitizeFileName(book.Author))
		err := os.MkdirAll(authorDir, 0755)
		if err != nil {
			return err
		}
		originalPath := filepath.Join(libPath, book.FileName)
		newPath := filepath.Join(authorDir, book.FileName)
		err = os.Rename(originalPath, newPath)
		if err != nil {
			return err
		}

	}
	return nil
}

func SanitizeFileName(name string) string {
	// platform specific
	// left as exercise
	return name
}
