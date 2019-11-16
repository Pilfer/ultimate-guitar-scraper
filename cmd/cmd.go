package cmd

import "github.com/urfave/cli"

// BeforeCommand sets the global flags before any commands are run
func BeforeCommand(c *cli.Context) error {
	return nil
}

// AfterCommand sets the global flags before any commands are run
func AfterCommand(c *cli.Context) error {
	return nil
}
