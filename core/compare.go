// core/compare.go
package core

type DiffResult struct {
	Status string // "added", "modified", "deleted"
	Path   string
}

// CompareSnapshots compares two file lists and returns a list of differences
func CompareSnapshots(current, reference []FileInfo) []DiffResult {
	refMap := make(map[string]string)
	for _, file := range reference {
		refMap[file.Path] = file.Hash
	}

	curMap := make(map[string]string)
	for _, file := range current {
		curMap[file.Path] = file.Hash
	}

	var diffs []DiffResult

	// Check for added and modified files
	for path, curHash := range curMap {
		if refHash, exists := refMap[path]; !exists {
			diffs = append(diffs, DiffResult{"added", path})
		} else if refHash != curHash {
			diffs = append(diffs, DiffResult{"modified", path})
		}
	}

	// Check for deleted files
	for path := range refMap {
		if _, exists := curMap[path]; !exists {
			diffs = append(diffs, DiffResult{"deleted", path})
		}
	}

	return diffs
}
