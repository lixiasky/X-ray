package core

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Scoreboard struct {
	counts map[string]int
	lock   sync.Mutex
}

var behaviorScores = &Scoreboard{
	counts: make(map[string]int),
}

// List of keywords considered high-risk file areas
var criticalPaths = []string{"/etc", "/usr", "/boot", "/bin", "/lib"}

const (
	ScoreThreshold = 5
	ScanInterval   = 5 * time.Second
)

// StartAutoDefense periodically evaluates behavior and triggers response
func StartAutoDefense() {
	for {
		analyzeAndRespond()
		time.Sleep(ScanInterval)
	}
}

func analyzeAndRespond() {
	events := GetBehaviorEvents()
	behaviorScores.lock.Lock()
	defer behaviorScores.lock.Unlock()

	for _, e := range events {
		// Accumulate score per source
		scoreKey := e.Source
		if scoreKey == "" {
			continue
		}

		// Increase score for critical write or many actions
		if strings.HasPrefix(e.EventType, "file_write") && isCriticalPath(e.Target) {
			behaviorScores.counts[scoreKey] += 3
		} else {
			behaviorScores.counts[scoreKey]++
		}

		// If score exceeds threshold, act!
		if behaviorScores.counts[scoreKey] >= ScoreThreshold {
			fmt.Printf("[autodefense] Threat detected from %s (score=%d). Eliminating...\n", scoreKey, behaviorScores.counts[scoreKey])

			KillProcessByName(scoreKey)
			DeleteFile(scoreKey)
			behaviorScores.counts[scoreKey] = -9999 // prevent repeat
			RecordEvent("auto_eliminate", scoreKey, "autodefense")
		}
	}
}

func isCriticalPath(path string) bool {
	for _, prefix := range criticalPaths {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}
	return false
}
