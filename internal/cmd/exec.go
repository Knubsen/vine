package cmd

import (
	"fmt"

	"github.com/Knubsen/vine/internal/utils"
	parser "github.com/Knubsen/vine/internal/yml_parser"
	"github.com/spf13/cobra"
)

var alias string

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "runs the thing",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		alias = args[0]

		config, err := parser.GetConfig()
		if err != nil {
			fmt.Printf("Error: %e", err)
			return
		}

		for _, c := range config.Commands {
			if utils.Contains(c.Aliases, alias) {
				fmt.Printf(c.Command)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
