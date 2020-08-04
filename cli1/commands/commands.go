package commands

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var (
	// RootCmd defines root command
	RootCmd = &cobra.Command {
		Use: "foobar",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
)

// Run runs command.
func Run() {
	RootCmd.Execute()
}

// Exit finishes a running action.
func Exit(err error, codes ...int) {
	var code intif len(codes) > 0 {
		code = codes[0]
	} else {
		code = 2
	}
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}