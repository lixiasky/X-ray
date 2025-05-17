// core/tracker.go
// Tracks behavior events and stores them in a chain for forensic analysis

package core

import (
	"fmt"
	"sync"
	"time"
)

// BehaviorEvent represents a single system behavior event
type BehaviorEvent struct {
	Timestamp time.Time // When the event occurred
	EventType string    // "file_create", "file_write", "file_delete", "process_start", etc.
	Target    string    // File path or process name
	Source    string    // Optional: who/what triggered it (e.g., parent process)
}

// BehaviorChain holds a list of behavior events
type BehaviorChain struct {
	events []BehaviorEvent
	lock   sync.Mutex
}

// Global behavior chain instance
var chain = &BehaviorChain{
	events: []BehaviorEvent{},
}

// RecordEvent appends a new behavior event to the global behavior chain
func RecordEvent(eventType, target, source string) {
	chain.lock.Lock()
	defer chain.lock.Unlock()

	event := BehaviorEvent{
		Timestamp: time.Now(),
		EventType: eventType,
		Target:    target,
		Source:    source,
	}
	chain.events = append(chain.events, event)
}

// PrintChain outputs the full behavior chain to console
func PrintChain() {
	chain.lock.Lock()
	defer chain.lock.Unlock()

	fmt.Println("[tracker] Behavior Chain:")
	for _, e := range chain.events {
		fmt.Printf(" - [%s] %s -> %s (source: %s)\n",
			e.Timestamp.Format(time.RFC3339),
			e.EventType,
			e.Target,
			e.Source,
		)
	}
}

// tracker.go
func GetBehaviorEvents() []BehaviorEvent {
	chain.lock.Lock()
	defer chain.lock.Unlock()
	return chain.events
}
