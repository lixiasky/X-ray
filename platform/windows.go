// windows.go
// Platform-specific implementation for Windows

//go:build windows

package platform

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// GetPlatformName returns the current platform
func GetPlatformName() string {
	return "Windows"
}

// KillProcessByName kills processes by image name (e.g., "notepad.exe") using taskkill
func KillProcessByName(name string) error {
	if !strings.HasSuffix(name, ".exe") {
		name += ".exe"
	}
	cmd := exec.Command("taskkill", "/IM", name, "/F")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("taskkill failed: %v\nOutput: %s", err, output)
	}
	return nil
}

// NormalizePath converts Linux-style paths to Windows equivalents
func NormalizePath(path string) string {
	if runtime.GOOS != "windows" {
		return path
	}
	if strings.HasPrefix(path, "/tmp") {
		return filepath.Join("C:\\Temp", strings.TrimPrefix(path, "/tmp"))
	} else if strings.HasPrefix(path, "/home") {
		return filepath.Join("C:\\Users\\Public", strings.TrimPrefix(path, "/home"))
	} else if strings.HasPrefix(path, "/etc") {
		return filepath.Join("C:\\ProgramData", strings.TrimPrefix(path, "/etc"))
	}
	return filepath.Join("C:\\XrayUnknown", strings.ReplaceAll(path, "/", "_"))
}

// MountISO is not supported on Windows in this tool (simulated failure)
func MountISO(isoPath string, mountPoint string) error {
	return fmt.Errorf("ISO mounting not supported on Windows")
}

// UnmountISO is a no-op on Windows
func UnmountISO(mountPoint string) error {
	return nil
}
