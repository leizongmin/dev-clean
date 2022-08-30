package diskutil

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/leizongmin/dev-clean/logutil"
)

func TestScan(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	assert.NoError(t, err)

	// for macOS
	skipRe, err := regexp.Compile(fmt.Sprintf(`^%s/(\..*|go|Library|Applications|Desktop|Documents|Mail|Movies|Music|Pictures|News)$`, homeDir))
	assert.NoError(t, err)
	assert.NotNil(t, skipRe)

	count, err := Scan(homeDir, func(path string, info os.FileInfo) bool {
		if skipRe.MatchString(path) {
			logutil.Warnf("skip: %s", path)
			return false
		}

		// logutil.Infof("add: %s", path)
		return true
	}, nil)
	assert.NoError(t, err)
	assert.Greater(t, count, int32(0))
	fmt.Println(count)
}

func TestScanDir(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	assert.NoError(t, err)
	re, err := regexp.Compile(`/(node_modules|target)$`)
	assert.NoError(t, err)
	assert.NotNil(t, re)

	// for macOS
	skipRe, err := regexp.Compile(fmt.Sprintf(`^%s/(\..*|go|Library|Applications|Desktop|Documents|Mail|Movies|Music|Pictures|News)$`, homeDir))
	assert.NoError(t, err)
	assert.NotNil(t, skipRe)

	var dirs []string
	count, err := ScanDir(homeDir, func(path string, info os.FileInfo) bool {
		if skipRe.MatchString(path) {
			logutil.Warnf("skip: %s", path)
			return false
		}

		if d, _ := os.Stat(filepath.Join(path, "node_modules")); d != nil && d.IsDir() {
			logutil.Infof("fast add: %s", path)
			dirs = append(dirs, path)
			return false
		}
		if d, _ := os.Stat(filepath.Join(path, "target")); d != nil && d.IsDir() {
			logutil.Infof("fast add: %s", path)
			dirs = append(dirs, path)
			return false
		}

		// if re.MatchString(path) {
		// 	logutil.Infof("add: %s", path)
		// 	dirs = append(dirs, path)
		// 	return false
		// }

		return true
	}, func(path string, err error) bool {
		return true
	})

	assert.NoError(t, err)
	assert.Greater(t, count, int32(0))
	fmt.Println(count)

	fmt.Println(dirs)
}
