package cli

import (
	"github.com/urfave/cli/v2"
)

// Command is a wrapper of cli.Command
type Command struct {
	cli.Command
}

// AddSubcommands adds subcommands
func (cmd *Command) AddSubcommands(commands ...*Command) {
	for _, command := range commands {
		cmd.Subcommands = append(cmd.Subcommands, &command.Command)
	}
}

// AddFlags adds flags
func (cmd *Command) AddFlags(flags ...*Flag) {
	for _, flag := range flags {
		cmd.Flags = append(cmd.Flags, flag)
	}
}

// SetBeforeFunc sets the Before fucntion for the cli.Command
func (cmd *Command) SetBeforeFunc(before func() error) {
	cmd.Before = func(c *cli.Context) error {
		return before()
	}
}

// SetActionFunc sets the Action fucntion for the cli.Command
func (cmd *Command) SetActionFunc(actionFn ActionFn) {
	cmd.Action = func(c *cli.Context) error {
		args := []string(c.Args().Tail())
		return actionFn(c.Context, args)
	}
}

// NewCommand create a new instance of the Command struct
func NewCommand() *Command {
	return &Command{
		Command: cli.Command{},
	}
}
