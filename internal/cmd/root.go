package cmd

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
	"github.com/savioxavier/termlink"
	"github.com/spf13/cobra"
)

var uiSize int

// Cli cmd struct
type Cli struct {
	rootCmd *cobra.Command
}

// NewCli returns the cli instance used to register and execute command
func NewCli() *Cli {
	cli := &Cli{
		rootCmd: &cobra.Command{
			Use:   "mchat",
			Short: "Enjoy the mini game.",
			Long:  printLogo(),
		},
	}
	cli.rootCmd.SetOut(os.Stdout)
	cli.rootCmd.SetErr(os.Stderr)
	cli.setFlags()
	cli.rootCmd.DisableAutoGenTag = true
	return cli
}

func (cli *Cli) setFlags() {
	flags := cli.rootCmd.PersistentFlags()
	flags.IntVar(&uiSize, "ui-size", 4, "number of list items to show in menu at once")
}

var Logo = "\n  __  __  ____                      \n |  \\/  |/ ___| __ _ _ __ ___   ___ \n | |\\/| | |  _ / _` | '_ ` _ \\ / _ \\\n | |  | | |_| | (_| | | | | | |  __/\n |_|  |_|\\____|\\__,_|_| |_| |_|\\___|\n                                    \n"

func printLogo() string {
	panel := pterm.DefaultHeader.WithMargin(8).
		WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).
		WithTextStyle(pterm.NewStyle(pterm.FgLightWhite)).Sprint("Enjoy the mini game.")
	logo := pterm.FgLightGreen.Sprint(Logo)
	pterm.Info.Prefix = pterm.Prefix{
		Text:  "Tips",
		Style: pterm.NewStyle(pterm.BgBlue, pterm.FgLightWhite),
	}
	url := pterm.Info.Sprintf("Find more information at: %s", termlink.ColorLink("mchat", "https://github.com/lcr2000/mchat", "italic green"))
	return fmt.Sprintf(`
%s%s
%s
`, panel, logo, url)
}
