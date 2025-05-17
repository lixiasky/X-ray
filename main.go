// main.go
// Entry point of X-RAY: Debug-enhanced mode with verbose outputs

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"xray/core"
	"xray/decoy"
	"xray/engine"
	"xray/export"
	"xray/platform"
)

func main() {
	fmt.Printf("[X-RAY] Starting on platform: %s\n", platform.GetPlatformName())

	// Step 1: Ask for ISO path
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("[X-RAY] Enter full path to your ISO file: ")
	isoPath, _ := reader.ReadString('\n')
	isoPath = isoPath[:len(isoPath)-1] // Trim newline

	// Step 2: Mount ISO
	referenceMount, err := engine.MountISO(isoPath)
	if err != nil {
		fmt.Printf("[X-RAY] ERROR: Failed to prepare reference snapshot: %v\n", err)
		return
	}
	defer engine.UnmountAll()

	fmt.Println("[X-RAY] Snapshotting current system...")
	current, _ := core.ScanSystem("/")
	fmt.Println("[X-RAY] Snapshotting ISO reference...")
	reference, _ := core.ScanSystem(referenceMount)
	diffs := core.CompareSnapshots(current, reference)
	fmt.Printf("[X-RAY] Found %d differences from reference.\n", len(diffs))

	fmt.Println("[X-RAY] Exporting logs...")
	export.ExportAsJSON("./behavior.json")
	export.ExportAsText("./behavior.log")
	fmt.Println("[X-RAY] Initial logs written.")

	fmt.Println("[X-RAY] Starting background systems...")
	go func() {
		fmt.Println("[system] MonitorEntireSystem started")
		core.MonitorEntireSystem()
	}()
	go func() {
		fmt.Println("[system] WatchProcesses started")
		core.WatchProcesses(3)
	}()
	go func() {
		fmt.Println("[system] StartAutoDefense started")
		core.StartAutoDefense()
	}()
	go func() {
		fmt.Println("[system] Decoy system started")
		decoy.StartDecoySystem()
	}()
	go func() {
		for {
			export.ExportGraphvizDOT("./behavior.dot")
			fmt.Println("[export] behavior.dot updated.")
			time.Sleep(30 * time.Second)
		}
	}()

	fmt.Println("[X-RAY] All systems running. Awaiting threats... Press Ctrl+C to terminate.")
	select {}
}
