// export/graph.go
// Exports the tracked behavior chain as a Graphviz DOT file

package export

import (
	"fmt"
	"os"
	"xray/core"
)

// ExportGraphvizDOT exports the global behavior chain as a .dot file for visualization
func ExportGraphvizDOT(filepath string) error {
	events := core.GetBehaviorEvents()

	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create graph file: %v", err)
	}
	defer file.Close()

	_, _ = file.WriteString("digraph BehaviorGraph {\n")
	_, _ = file.WriteString("  rankdir=LR;\n")
	_, _ = file.WriteString("  node [shape=box, fontname=\"Courier\"];\n")

	for i, e := range events {
		nodeID := fmt.Sprintf("e%d", i)
		label := fmt.Sprintf("%s\\n%s", e.EventType, e.Target)
		_, _ = file.WriteString(fmt.Sprintf("  %s [label=\"%s\"];\n", nodeID, label))

		if e.Source != "" {
			for j, s := range events {
				if s.Target == e.Source {
					_, _ = file.WriteString(fmt.Sprintf("  e%d -> %s;\n", j, nodeID))
					break
				}
			}
		}
	}

	_, _ = file.WriteString("}\n")
	fmt.Printf("[graph] Exported behavior chain to %s\n", filepath)
	return nil
}
