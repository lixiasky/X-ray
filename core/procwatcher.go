// core/procwatcher.go
// Monitors /proc for new process creation by checking for new PIDs

package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var knownPIDs = make(map[int]bool)

// WatchProcesses continuously checks for new PIDs in /proc
func WatchProcesses(intervalSec int) {
	for {
		entries, _ := os.ReadDir("/proc")

		for _, e := range entries {
			pid, err := strconv.Atoi(e.Name())
			if err != nil {
				continue // not a PID directory
			}

			if !knownPIDs[pid] {
				knownPIDs[pid] = true
				path := filepath.Join("/proc", e.Name(), "exe")
				target, err := os.Readlink(path)
				if err == nil {
					fmt.Printf("[proc] New process: %d -> %s\n", pid, target)
					RecordEvent("process_start", target, "proc_watcher")
				}
			}
		}

		time.Sleep(time.Duration(intervalSec) * time.Second)
	}
}
