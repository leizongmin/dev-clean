package diskutil

import (
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

func Scan(dir string, onEntry func(path string, info os.FileInfo) bool, onError func(path string, err error) bool) (count int32, err error) {
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

	wg := sync.WaitGroup{}
	end := atomic.Bool{}

	for _, entry := range list {
		if end.Load() {
			break
		}
		count++

		path := filepath.Join(dir, entry.Name())
		info, err := entry.Info()
		if err != nil {
			if !onError(path, err) {
				end.Store(true)
			}
		}

		if !onEntry(path, info) {
			continue
		}

		wg.Add(1)
		go func(entry os.DirEntry) {
			defer wg.Done()
			if info.IsDir() {
				n, err := Scan(path, onEntry, onError)
				if err != nil {
					if !onError(path, err) {
						end.Store(true)
					}
				}
				atomic.AddInt32(&count, n)
			}
		}(entry)
	}

	wg.Wait()

	return count, nil
}

func ScanDir(dir string, onEntry func(path string, info os.FileInfo) bool, onError func(path string, err error) bool) (count int32, err error) {
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

	wg := sync.WaitGroup{}
	end := atomic.Bool{}

	for _, entry := range list {
		if end.Load() {
			break
		}
		count++
		if entry.IsDir() {
			wg.Add(1)
			go func(entry os.DirEntry) {
				defer wg.Done()

				path := filepath.Join(dir, entry.Name())
				info, err := entry.Info()
				if err != nil {
					if !onError(path, err) {
						end.Store(true)
					}
					return
				}

				if !onEntry(path, info) {
					return
				}

				n, err := ScanDir(path, onEntry, onError)
				if err != nil {
					if !onError(path, err) {
						end.Store(true)
					}
				}
				atomic.AddInt32(&count, n)
			}(entry)
		}
	}

	wg.Wait()

	return count, nil
}
