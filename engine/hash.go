// engine/hash.go

package engine

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// ComputeHash computes SHA256 hash of a file and returns it as hex string
func ComputeHash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file %s: %v", path, err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("failed to hash file %s: %v", path, err)
	}

	hashSum := hasher.Sum(nil)
	return hex.EncodeToString(hashSum), nil
}
