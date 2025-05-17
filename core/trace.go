// trace.go
// Traces suspicious behavior chains back to source and eliminates all related components

package core

import (
	"fmt"
	"log"
	"strings"
)

// TraceAndPurge searches the behavior chain for any event involving the target,
// then backtracks to identify its origin and eliminate the full path
func TraceAndPurge(target string) {
	events := GetBehaviorEvents()
	var source string = ""

	// Step 1: find who wrote or launched the target
	for _, e := range events {
		if strings.Contains(e.Target, target) {
			source = e.Source
			break
		}
	}

	if source == "" {
		log.Printf("[trace] No source found for %s. Skipping.", target)
		return
	}

	log.Printf("[trace] Tracing origin of %s -> %s", target, source)

	// Step 2: eliminate the source file and process
	err1 := DeleteFile(source)
	err2 := KillProcessByName(source)
	RecordEvent("trace_cleanup", source, "trace")

	// Step 3: cascade purge any events caused by the source
	for _, e := range events {
		if e.Source == source {
			DeleteFile(e.Target)
			KillProcessByName(e.Target)
			RecordEvent("trace_cascade", e.Target, source)
		}
	}

	fmt.Printf("[trace] Cleanup complete for chain originating from %s\n", source)

	if err1 != nil || err2 != nil {
		log.Printf("[trace] Warnings during purge: %v | %v", err1, err2)
	}
}
