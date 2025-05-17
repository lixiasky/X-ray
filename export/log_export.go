// export/log_export.go
// Exports the behavior chain as plain text and JSON logs for archival or external analysis

package export

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"xray/core"
)

// ExportAsJSON writes the behavior events into a JSON file
func ExportAsJSON(filename string) error {
	events := core.GetBehaviorEvents()

	data, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON file: %v", err)
	}

	fmt.Printf("[log] Exported behavior log to %s\n", filename)
	return nil
}

// ExportAsText writes a human-readable log file
func ExportAsText(filename string) error {
	events := core.GetBehaviorEvents()

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create log file: %v", err)
	}
	defer file.Close()

	for _, e := range events {
		line := fmt.Sprintf("[%s] %s: %s (source: %s)\n",
			e.Timestamp.Format(time.RFC3339),
			e.EventType,
			e.Target,
			e.Source,
		)
		_, _ = file.WriteString(line)
	}

	fmt.Printf("[log] Exported readable log to %s\n", filename)
	return nil
}

// ExportBehaviorJSON writes the full behavior chain as JSON array to behavior.json
func ExportBehaviorJSON(path string) error {
	events := core.GetBehaviorEvents()
	data, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
