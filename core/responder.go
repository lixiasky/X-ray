// core/responder.go
// Provides automated responses to suspicious or malicious behavior

package core

import (
	"fmt"
	"os"
	"os/exec"
)

// KillProcess attempts to terminate a process by PID
func KillProcess(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("cannot find process %d: %v", pid, err)
	}

	err = process.Kill()
	if err != nil {
		return fmt.Errorf("failed to kill process %d: %v", pid, err)
	}

	fmt.Printf("[responder] Killed process PID %d\n", pid)
	return nil
}

// KillProcessByName kills all processes matching the given name using pkill
func KillProcessByName(name string) error {
	cmd := exec.Command("pkill", "-f", name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("pkill failed for %s: %v - %s", name, err, string(output))
	}

	fmt.Printf("[responder] Killed all processes named %s\n", name)
	return nil
}

// DeleteFile attempts to delete a file at the given path
func DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return fmt.Errorf("failed to delete file %s: %v", path, err)
	}

	fmt.Printf("[responder] Deleted file: %s\n", path)
	return nil
}
