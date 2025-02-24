package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "runs the thing",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Execute command with alias: ", args)
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
