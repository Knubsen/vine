package runauthorization

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Knubsen/vine/internal/models"
	"github.com/spf13/cobra"
)

func AuthRun(fn func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {

		// authorization
		fmt.Println(cmd.CommandPath())
		fmt.Println(cmd.CalledAs())

		authorizeCommand(cmd)

		fmt.Println(os.Args[1:])
		fmt.Println(os.Executable())
		fmt.Println(args)

		fn(cmd, args)
	}
}

func authorizeCommand(cmd *cobra.Command) (string, bool) {
	commandConfig := models.CommandAuthConfig{}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	jsonstr, err := os.ReadFile(filepath.Join(path, "/config/meeh.json"))
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	err = json.Unmarshal([]byte(jsonstr), &commandConfig)
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	fmt.Println(commandConfig)

	return "", false
}
