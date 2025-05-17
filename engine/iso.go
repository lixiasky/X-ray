package engine

import (
	//"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

var (
	IsoMountPoint      = "/mnt/xray_iso"
	SquashfsMountPoint = "/mnt/xray_root"
)

// MountISO mounts the ISO file and returns the path to the mounted squashfs
func MountISO(isoPath string) (string, error) {
	fmt.Println("[ISO] Mounting ISO file:", isoPath)

	// Create mount points
	if err := os.MkdirAll(IsoMountPoint, 0755); err != nil {
		return "", fmt.Errorf("failed to create ISO mount point: %v", err)
	}
	if err := os.MkdirAll(SquashfsMountPoint, 0755); err != nil {
		return "", fmt.Errorf("failed to create squashfs mount point: %v", err)
	}

	// Mount the ISO
	cmd := exec.Command("mount", "-o", "loop", isoPath, IsoMountPoint)
	output, err := cmd.CombinedOutput()
	if err != nil && !strings.Contains(string(output), "already mounted") {
		return "", fmt.Errorf("failed to mount ISO: %v\nOutput: %s", err, string(output))
	}

	// Find all .squashfs files and sort by size (desc)
	var squashfsFiles []string
	err = filepath.Walk(IsoMountPoint, func(path string, info os.FileInfo, err error) error {
		if err != nil || info == nil {
			return nil
		}
		if strings.HasSuffix(info.Name(), ".squashfs") {
			squashfsFiles = append(squashfsFiles, path)
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("error while searching for squashfs: %v", err)
	}
	if len(squashfsFiles) == 0 {
		return "", fmt.Errorf("no .squashfs file found in ISO — please ensure this is a live system ISO like Kali or Ubuntu")
	}

	// Sort and pick the largest squashfs file
	sort.Slice(squashfsFiles, func(i, j int) bool {
		fi1, _ := os.Stat(squashfsFiles[i])
		fi2, _ := os.Stat(squashfsFiles[j])
		return fi1.Size() > fi2.Size()
	})
	squashfsFile := squashfsFiles[0]
	fmt.Println("[ISO] Using squashfs file:", squashfsFile)

	// Mount the squashfs
	cmd = exec.Command("mount", "-t", "squashfs", "-o", "loop", squashfsFile, SquashfsMountPoint)
	output, err = cmd.CombinedOutput()
	if err != nil && !strings.Contains(string(output), "already mounted") {
		return "", fmt.Errorf("failed to mount squashfs: %v\nOutput: %s", err, string(output))
	}

	fmt.Println("[ISO] ISO and squashfs mounted successfully.")
	return SquashfsMountPoint, nil
}

// UnmountAll cleans up mount points
func UnmountAll() {
	fmt.Println("[ISO] Unmounting all...")

	if _, err := os.Stat(SquashfsMountPoint); err == nil {
		exec.Command("umount", SquashfsMountPoint).Run()
	}
	if _, err := os.Stat(IsoMountPoint); err == nil {
		exec.Command("umount", IsoMountPoint).Run()
	}
}
