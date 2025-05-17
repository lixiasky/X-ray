// core/scanner.go
package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"xray/engine"
)

// FileInfo holds path and hash of a scanned file
type FileInfo struct {
	Path string
	Hash string
}

// ScanSystem scans all regular files under rootPath and returns their hash snapshots
func ScanSystem(rootPath string) ([]FileInfo, error) {
	var files []FileInfo
	var counter int

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("[scan] error accessing %s: %v\n", path, err)
			return nil
		}

		// Skip virtual/volatile system paths
		if strings.HasPrefix(path, "/proc") || strings.HasPrefix(path, "/sys") || strings.HasPrefix(path, "/dev") {
			return nil
		}

		if !info.Mode().IsRegular() {
			return nil // skip dirs, links, etc.
		}

		hash, err := engine.ComputeHash(path)
		if err != nil {
			fmt.Printf("[scan] failed to hash %s: %v\n", path, err)
			return nil
		}

		files = append(files, FileInfo{
			Path: path,
			Hash: hash,
		})

		counter++
		if counter%500 == 0 {
			fmt.Printf("[scan] %d files scanned...\n", counter)
		}

		return nil
	})

	fmt.Printf("[scan] Done. Total scanned: %d files.\n", counter)
	return files, err
}
