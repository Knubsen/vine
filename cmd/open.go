package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "runs the thing",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Opened", args)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
