package diskutil

import (
	"os"
	"path/filepath"
)

func Scan(dir string, onEntry func(path string, info os.FileInfo) bool, onError func(path string, err error) bool) (count int, err error) {
	dir, err = filepath.Abs(dir)
	if err != nil {
		return count, err
	}

	if onError == nil {
		onError = func(string, error) bool {
			return false
		}
	}

	list, err := os.ReadDir(dir)
	if err != nil {
		if onError(dir, err) {
			return count, nil
		}
		return count, err
	}

	for _, entry := range list {
		count++
		path := filepath.Join(dir, entry.Name())
		info, err := entry.Info()
		if err != nil {
			if onError(path, err) {
				continue
			}
			return count, err
		}

		if !onEntry(path, info) {
			continue
		}

		if info.IsDir() {
			n, err := Scan(path, onEntry, onError)
			if err != nil {
				if onError(path, err) {
					continue
				}
				return count, err
			}
			count += n
		}
	}

	return count, nil
}

func ScanDir(dir string, onEntry func(path string, info os.FileInfo) bool, onError func(path string, err error) bool) (count int, err error) {
	dir, err = filepath.Abs(dir)
	if err != nil {
		return count, err
	}

	if onError == nil {
		onError = func(string, error) bool {
			return false
		}
	}

	list, err := os.ReadDir(dir)
	if err != nil {
		if onError(dir, err) {
			return count, nil
		}
		return count, err
	}

	for _, entry := range list {
		count++
		if entry.IsDir() {
			path := filepath.Join(dir, entry.Name())
			info, err := entry.Info()
			if err != nil {
				if onError(path, err) {
					continue
				}
				return count, err
			}

			if !onEntry(path, info) {
				continue
			}

			n, err := Scan(path, onEntry, onError)
			if err != nil {
				if onError(path, err) {
					continue
				}
				return count, err
			}
			count += n
		}
	}

	return count, nil
}
