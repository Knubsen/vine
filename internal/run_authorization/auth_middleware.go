package runauthorization

import "github.com/spf13/cobra"

func AuthRun(fn func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	// return func(cmd *cobra.Command, args []string) {
	//
	// authorization
	//
	// 	fn(cmd, args)
	// }
}
