package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vine",
	Short: "quick access stuff",
	Run: func(cmd *cobra.Command, args []string) {
		// openCmd.Run(cmd, args)

		dFlag, _ := cmd.Flags().GetString("d")
		if dFlag != "" {
			fmt.Printf("Flag -d reveived with input %s\n", dFlag)
		} else {
			fmt.Println(" No flag -d provided")
		}
	},
}

func Execute() {
	rootCmd.Flags().StringP("d", "d", "", "Some text here")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
