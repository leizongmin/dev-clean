package terminalutil

import (
	"os"

	"github.com/cheggaaa/pb/v3"
	"github.com/manifoldco/promptui"

	"github.com/leizongmin/dev-clean/logutil"
)

func Confirm(message string) bool {
	prompt := promptui.Select{
		Label: message,
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()

	if err != nil {
		logutil.Fatalf("%s", err)
		os.Exit(1)
	}

	return result == "Yes"
}

func Progress(prefix string, start int, total int, f func(add func(int), done func())) {
	bar := pb.New(total)
	bar.Add(start)
	bar.SetTemplateString(`{{with string . "prefix"}}{{.}} {{end}}{{counters . }} {{bar . }} {{percent . }} {{with string . "suffix"}} {{.}}{{end}}`)
	bar.Set("prefix", prefix)
	bar.Set("speed", false)
	bar.Start()

	add := func(v int) {
		bar.Add(v)
	}

	done := func() {
		bar.Finish()
	}

	f(add, done)
}
