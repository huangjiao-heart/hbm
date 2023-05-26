package main

import (
	"fmt"

	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/cli/command/commands"
	"github.com/kassisol/hbm/pkg/juliengk/go-utils/filedir"
)

func main() {
	scPath := "/tmp/hbm/shellcompletion"
	bashTarget := fmt.Sprintf("%s/bash", scPath)

	if err := filedir.CreateDirIfNotExist(scPath, true, 0755); err != nil {
		fmt.Println(err)
	}

	cmd := command.NewHBMCommand()
	commands.AddCommands(cmd)
	cmd.DisableAutoGenTag = true

	if err := cmd.GenBashCompletionFile(bashTarget); err != nil {
		fmt.Println(err)
	}
}
