package commands

import (
	"strings"
	"time"

	"github.com/abiosoft/ishell"
)

func Edit(shell *ishell.Shell) (f func(c *ishell.Context)) {
	return func(c *ishell.Context) {
		if errString, ok := syntaxCheck(c, 1); !ok {
			shell.Println(errString)
			return
		}
		path := strings.Join(c.Args, " ")
		entry, ok := getEntryByPath(shell, path)
		if !ok {
			shell.Printf("couldn't find entry '%s'\n", path)
			return
		}
		shell.ShowPrompt(false)
		if err := promptForEntry(shell, entry, entry.Title()); err != nil {
			shell.Printf("couldn't edit entry: %s\n", err)
		}
		entry.SetLastModificationTime(time.Now())

		shell.ShowPrompt(true)
	}
}
