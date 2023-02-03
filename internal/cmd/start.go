package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/lcr2000/mchat/internal/server"
	"github.com/spf13/cobra"
)

// StartCommand start cmd struct
type StartCommand struct {
	BaseCommand
}

// Init ListCommand
func (c *StartCommand) Init() {
	c.command = &cobra.Command{
		Use:     "start",
		Short:   "Start the server",
		Long:    "Start the server",
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.runStart(cmd, args)
		},
		Example: startExample(),
	}
	c.command.DisableFlagsInUseLine = true
}

func (c *StartCommand) runStart(command *cobra.Command, args []string) error {
	server.StartServer()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // Read user input.
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // Exit if enter Q.
			return nil
		}
	}
}

func startExample() string {
	return `
# Start the server
mchat start
# Aliases
mchat s
`
}
