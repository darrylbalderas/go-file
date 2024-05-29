package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/sys/unix"
)

func printUnixStat(path string) error {
	var stat unix.Stat_t
	if err := unix.Stat(path, &stat); err != nil {
		return err
	}

	atime := time.Unix(stat.Atim.Sec, stat.Atim.Nsec)
	mtime := time.Unix(stat.Mtim.Sec, stat.Mtim.Nsec)
	ctime := time.Unix(stat.Ctim.Sec, stat.Ctim.Nsec)

	fmt.Printf("File: %s\n", path)
	fmt.Printf("Device: %d\n", stat.Dev)
	fmt.Printf("Inode: %d\n", stat.Ino)
	fmt.Printf("Mode: %o\n", stat.Mode)
	fmt.Printf("Link count: %d\n", stat.Nlink)
	fmt.Printf("UID: %d\n", stat.Uid)
	fmt.Printf("GID: %d\n", stat.Gid)
	fmt.Printf("Rdev: %d\n", stat.Rdev)
	fmt.Printf("Size: %d bytes\n", stat.Size)
	fmt.Printf("Block size: %d bytes\n", stat.Blksize)
	fmt.Printf("Blocks: %d\n", stat.Blocks)
	fmt.Printf("Access time: %s\n", atime)
	fmt.Printf("Modify time: %s\n", mtime)
	fmt.Printf("Change time: %s\n", ctime)

	return nil
}

func getAccessTime(path string) (time.Time, error) {
	var stat unix.Stat_t
	if err := unix.Stat(path, &stat); err != nil {
		return time.Time{}, err
	}

	atime := time.Unix(stat.Atim.Sec, stat.Atim.Nsec)
	fmt.Printf("AccessTime: %s\n", atime)
	return atime, nil
}

func printFileInfo(path string, info os.FileInfo) {
	fmt.Printf("Path: %s\n", path)
	fmt.Printf("Name: %s\n", info.Name())
	fmt.Printf("Size: %d bytes\n", info.Size())
	fmt.Printf("IsDir: %t\n", info.IsDir())
	fmt.Printf("Mode: %s\n", info.Mode())
	fmt.Printf("ModTime: %s\n", info.ModTime())
	fmt.Printf("IsRegular: %t\n", info.Mode().IsRegular())
}

func main() {
	// Call the function
	log.Println("Hello, World!")
	root := "./folders"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fileinfo, err := os.Stat(path)
		if err != nil {
			fmt.Printf("Error getting the file info: %v\n", err)
		}

		if info.IsDir() {
			fmt.Printf("Directory: %s\n", path)
		} else {
			fmt.Printf("File: %s\n", path)
		}

		printFileInfo(path, fileinfo)
		getAccessTime(path)
		printUnixStat(path)
		fmt.Println("")
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", root, err)
	}
}
