// linux.go
// Platform-specific implementation for Linux or macOS

//go:build linux || darwin

package platform

import (
	"fmt"
	"os/exec"

	//"path/filepath"
	"runtime"
)

// GetPlatformName returns the OS name
func GetPlatformName() string {
	return runtime.GOOS
}

// KillProcessByName uses pkill to terminate processes by name
func KillProcessByName(name string) error {
	cmd := exec.Command("pkill", "-f", name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("pkill failed: %v - %s", err, string(output))
	}
	return nil
}

// NormalizePath returns path unchanged on Linux/macOS
func NormalizePath(path string) string {
	return path
}

// MountISO mounts an ISO file using loop device
func MountISO(isoPath string, mountPoint string) error {
	cmd := exec.Command("sudo", "mount", "-o", "loop", isoPath, mountPoint)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("mount failed: %v - %s", err, string(output))
	}
	return nil
}

// UnmountISO unmounts the ISO mount point
func UnmountISO(mountPoint string) error {
	cmd := exec.Command("sudo", "umount", mountPoint)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("umount failed: %v - %s", err, string(output))
	}
	return nil
}
