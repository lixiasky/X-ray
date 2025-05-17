// decoy.go
// Deploys fake files to bait malware and logs any access to them

package decoy

import (
	"log"
	"os"
	"xray/core"

	"github.com/fsnotify/fsnotify"
)

// List of decoy files to deploy
var DecoyFiles = []string{
	"/tmp/fake_token.txt",
	"/home/fake_ssh_key",
	"/tmp/.xray_decoy_config.ini",
}

// Deploys decoy files with dummy content
func DeployDecoys() error {
	for _, path := range DecoyFiles {
		err := os.WriteFile(path, []byte("[DECOY] DO NOT TOUCH"), 0644)
		if err != nil {
			log.Printf("[decoy] Failed to create %s: %v", path, err)
		} else {
			log.Printf("[decoy] Deployed decoy file: %s", path)
		}
	}
	return nil
}

// WatchDecoys monitors decoy files for any access/modification
func WatchDecoys() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Printf("[decoy] Access detected: %s %s", event.Op, event.Name)
				core.RecordEvent("decoy_access", event.Name, "decoy")
				core.DeleteFile(event.Name)       // Optional: auto delete
				core.KillProcessByName("unknown") // Placeholder: future PID tracing

			case err := <-watcher.Errors:
				log.Printf("[decoy] Watch error: %v", err)
			}
		}
	}()

	for _, path := range DecoyFiles {
		err := watcher.Add(path)
		if err != nil {
			log.Printf("[decoy] Failed to watch %s: %v", path, err)
		} else {
			log.Printf("[decoy] Watching %s", path)
		}
	}

	return nil
}

// StartDecoySystem deploys and monitors decoy files
func StartDecoySystem() {
	DeployDecoys()
	err := WatchDecoys()
	if err != nil {
		log.Printf("[decoy] Failed to start decoy watcher: %v", err)
	}
	log.Printf("[decoy] Decoy system active.")
	select {} // Keep running
}
