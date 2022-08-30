package main

import (
	"errors"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"

	"github.com/dustin/go-humanize"
	"github.com/urfave/cli/v2"

	"github.com/leizongmin/dev-clean/logutil"
	"github.com/leizongmin/dev-clean/terminalutil"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "verbose", Aliases: []string{"v"}, Value: false, Usage: "show debug log"},
		},
		ArgsUsage: "<dir1> <dir2> ...",
		Action:    start,
		Before: func(ctx *cli.Context) error {
			if ctx.Bool("verbose") {
				logutil.SetLevel(logutil.LevelDebug)
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		logutil.Errorf("%s", err)
		os.Exit(1)
	}
}

func start(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		return errors.New("missing <dir1> <dir2> ...")
	}

	var dirs []string
	for _, dir := range ctx.Args().Slice() {
		results, err := ScanProjectCacheDirs(dir)
		if err != nil {
			logutil.Errorf("%s", err)
			continue
		}
		dirs = append(dirs, results...)
	}

	var counter int32
	var projects []ProjectItem
	terminalutil.Progress("counting file size", 0, len(dirs), func(set func(int), done func()) {
		wg := sync.WaitGroup{}
		for _, dir := range dirs {
			wg.Add(1)
			go func(dir string) {
				defer func() {
					wg.Done()
					atomic.AddInt32(&counter, 1)
					set(int(counter))
				}()

				p := ProjectItem{
					Type: DetectProjectType(dir),
					Root: dir,
				}
				switch p.Type {
				case "node":
					p.CacheDir = filepath.Join(dir, "node_modules")
				case "rust":
					p.CacheDir = filepath.Join(dir, "target")
				default:
					p.CacheDir = p.Root
				}
				size, err := CountDirTotalSize(p.CacheDir)
				if err != nil {
					logutil.Errorf("%s: %s", dir, err)
					return
				}
				p.Size = size
				projects = append(projects, p)
			}(dir)
		}
		wg.Wait()
		done()
	})

	for _, p := range projects {
		logutil.Infof("size: %s\ttype: %s \tdir: %s", humanize.Bytes(uint64(p.Size)), p.Type, p.CacheDir)
	}

	return nil
}

type ProjectItem struct {
	Type     string
	Root     string
	CacheDir string
	Size     int64
}

func DetectProjectType(dir string) string {
	if hasFile(dir, "package.json") && hasDir(dir, "node_modules") {
		return "node"
	}
	if hasFile(dir, "Cargo.toml") && hasDir(dir, "target") {
		return "rust"
	}
	return "unknown"
}

func hasFile(dir string, name string) bool {
	info, err := os.Stat(filepath.Join(dir, name))
	return err == nil && info.Mode().IsRegular()
}

func hasDir(dir string, name string) bool {
	info, err := os.Stat(filepath.Join(dir, name))
	return err == nil && info.Mode().IsDir()
}
