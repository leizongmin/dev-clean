package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/leizongmin/dev-clean/terminalutil"
)

func main() {
	if !terminalutil.Confirm("Are you sure continue") {
		return
	}

	terminalutil.Progress("Scanning files...", 0, 10, func(add func(int), done func()) {
		for i := 0; i < 10; i++ {
			add(1)
			time.Sleep(time.Millisecond * 200)
		}
		done()
	})

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "X", Value: "GET", Usage: "请求方法"},
			&cli.BoolFlag{Name: "i", Value: false, Usage: "是否包含完整输出"},
			&cli.StringSliceFlag{Name: "H", Usage: "请求头"},
		},
		Action: func(c *cli.Context) error {
			fmt.Println(c.String("X"))
			fmt.Println(c.Bool("i"))
			fmt.Println(c.StringSlice("H"))
			fmt.Println(c.Args().First())
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
