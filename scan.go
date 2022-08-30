package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync/atomic"

	"github.com/leizongmin/dev-clean/diskutil"
	"github.com/leizongmin/dev-clean/logutil"
)

func ScanProjectCacheDirs(root string) (dirs []string, err error) {
	logutil.Infof("scan %s", root)
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	skipRe, err := regexp.Compile(fmt.Sprintf(`^%s/(\..*|go|Library|Applications|Desktop|Documents|Mail|Movies|Music|Pictures|News)$`, homeDir))
	if err != nil {
		return nil, err
	}

	count, err := diskutil.ScanDir(root, func(path string, info os.FileInfo) bool {
		if skipRe.MatchString(path) {
			logutil.Debugf("skip: %s", path)
			return false
		}

		// node_modules
		if d, _ := os.Stat(filepath.Join(path, "node_modules")); d != nil && d.IsDir() {
			dirs = append(dirs, path)
			logutil.Debugf("add: %s", path)
			return false
		}
		// target
		if d, _ := os.Stat(filepath.Join(path, "target")); d != nil && d.IsDir() {
			dirs = append(dirs, path)
			logutil.Debugf("add: %s", path)
			return false
		}

		return true
	}, func(path string, err error) bool {
		logutil.Warnf("%s: %s", path, err)
		return true
	})

	logutil.Infof("scan %d files, found %d project cache dirs", count, len(dirs))

	return dirs, err
}

func CountDirTotalSize(root string) (size int64, err error) {
	_, err = diskutil.Scan(root, func(path string, info os.FileInfo) bool {
		if !info.IsDir() {
			logutil.Debugf("file: %s size: %d", path, info.Size())
			atomic.AddInt64(&size, info.Size())
		}
		return true
	}, func(path string, err error) bool {
		logutil.Warnf("%s: %s", path, err)
		return true
	})
	return size, err
}
