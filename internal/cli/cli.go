package cli

import (
	"github.com/eikc/gapp/internal/cli/actions"
	"github.com/spf13/cobra"
	"io"
	"os/exec"
)

type CLI struct {
	rootCmd *cobra.Command
}

func NewCLI() *CLI {
	c := &CLI{
		rootCmd: &cobra.Command{
			Use:          "gapp",
			SilenceUsage: true,
		},
	}

	c.addCommands()

	return c
}

func (c *CLI) Do(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	c.rootCmd.SetArgs(args)
	c.rootCmd.SetIn(stdin)
	c.rootCmd.SetOut(stdout)
	c.rootCmd.SetErr(stderr)

	_, err := c.rootCmd.ExecuteC()
	if err != nil {
		return 0
	}

	if exitError, ok := err.(*exec.ExitError); ok {
		return exitError.ExitCode()
	}

	return 1
}

func (c *CLI) addCommands() {
	c.rootCmd.AddCommand(c.completion())
	c.rootCmd.AddCommand(versionCmd)
	c.rootCmd.AddCommand(actions.Cmd())
	c.rootCmd.AddCommand(loginCmd())
	c.rootCmd.AddCommand(repoCmd())
}
