// core/monitor.go
// Recursively watches all relevant system directories for file changes using fsnotify

package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

// Excluded paths that should NOT be watched (system virtual dirs)
var ExcludedDirs = []string{
	"/proc", "/sys", "/dev", "/run", "/tmp", "/var/run", "/mnt", "/media",
}

// isExcluded checks if a path should be skipped from watch
func isExcluded(path string) bool {
	for _, ex := range ExcludedDirs {
		if strings.HasPrefix(path, ex) {
			return true
		}
	}
	return false
}

// MonitorEntireSystem walks through filesystem and adds all valid directories to watcher
func MonitorEntireSystem() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create watcher: %v", err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Printf("[monitor] File Event: %s %s\n", event.Op, event.Name)
				RecordEvent("file_"+event.Op.String(), event.Name, "fsnotify")

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Printf("[monitor] Watch error: %v", err)
			}
		}
	}()

	walkErr := filepath.Walk("/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("[monitor] Access error at %s: %v", path, err)
			return nil
		}
		if !info.IsDir() || isExcluded(path) {
			return nil
		}
		if err := watcher.Add(path); err != nil {
			log.Printf("[monitor] Failed to watch %s: %v", path, err)
		}
		return nil
	})

	if walkErr != nil {
		log.Printf("[monitor] Walk error: %v", walkErr)
	}

	log.Println("[monitor] All applicable directories added.")
	select {}
}
