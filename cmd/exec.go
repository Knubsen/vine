package cmd

import (
	"fmt"

	parser "github.com/Knubsen/vine/yml_parser"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "runs the thing",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Execute command with alias: ", args)

		config, err := parser.GetConfig()
		if err != nil {
			fmt.Printf("Error: %e", err)
			return
		}

		for _, c := range config.Commands {
			fmt.Printf("Command: %s\n", c.Command)
			fmt.Printf("Aliases: %v\n", c.Aliases)
		}
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
