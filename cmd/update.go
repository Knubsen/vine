package cmd

import (
	"fmt"

	runauthorization "github.com/Knubsen/vine/internal/run_authorization"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "runs the thing",
	Run: runauthorization.AuthRun(func(cmd *cobra.Command, args []string) {
		fmt.Println("Authorized run")
	}),
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
