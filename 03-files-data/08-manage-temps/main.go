package main

import "fmt"

func main() {
	tempFile, err := CreateTempFile("librago")
	if err != nil {
		fmt.Printf("Failed to create a temp file: %s\n", err)
		return
	}
	ProcessAndCleanupTempFile(tempFile)
	tempDir, err := CreateTempDir("librago")
	if err != nil {
		fmt.Printf("Failed to create a temp dir: %s\n", err)
		return
	}
	ProcessAndCleanupTempDir(tempDir)

}
