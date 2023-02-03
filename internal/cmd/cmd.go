package cmd

// NewBaseCommand cmd struct
func NewBaseCommand() *BaseCommand {
	cli := NewCli()
	baseCmd := &BaseCommand{
		command: cli.rootCmd,
	}
	baseCmd.AddCommands(
		&StartCommand{},
		&ChatCommand{},
		&ListOnlineCommand{},
	)

	return baseCmd
}
