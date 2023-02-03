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

// Init StartCommand
func (c *StartCommand) Init() {
	c.command = &cobra.Command{
		Use:     "start",
		Short:   "Start the server",
		Long:    "Start the server",
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.run()
		},
		Example: c.example(),
	}
	c.command.DisableFlagsInUseLine = true
}

func (c *StartCommand) run() error {
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

func (c *StartCommand) example() string {
	return `
# Start the server
mchat start
# Aliases
mchat s
`
}
