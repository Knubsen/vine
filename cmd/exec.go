package cmd

import (
	"fmt"

	runauthorization "github.com/Knubsen/vine/internal/run_authorization"
	"github.com/Knubsen/vine/internal/utils"
	parser "github.com/Knubsen/vine/internal/yml_parser"
	"github.com/spf13/cobra"
)

var alias string

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "runs the thing",
	Args:  cobra.ExactArgs(1),
	Run: runauthorization.AuthRun(
		func(cmd *cobra.Command, args []string) {
			alias = args[0]

			vault, err := parser.GetVault()
			if err != nil {
				fmt.Printf("Error: %e", err)
				return
			}

			for _, c := range vault.Commands {
				if utils.Contains(c.Aliases, alias) {
					fmt.Printf(c.Command)
					return
				}
			}
		}),
}

func init() {
	rootCmd.AddCommand(execCmd)
}
