package runauthorization

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/Knubsen/vine/internal/models"
	"github.com/Knubsen/vine/internal/state"
	"github.com/Knubsen/vine/internal/utils"
	"github.com/spf13/cobra"
)

func AuthRun(fn func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		allowed, err := authorizeCommand(cmd)
		if err != nil {
			fmt.Printf("Error authorizing command: %s\n", err)
			os.Exit(1)
		}

		if !allowed {
			fmt.Printf("You do not have permission to run this command: %s/n", cmd.CalledAs())
			os.Exit(1)
		}

		fn(cmd, args)
	}
}

func authorizeCommand(cmd *cobra.Command) (bool, error) {
	vinerunAuths := path.Join(state.Location.Auth, ".vine_run_auth.json")
	commands := models.AuthorizedCommand{}

	jsonstr, err := os.ReadFile(vinerunAuths)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal([]byte(jsonstr), &commands)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	permission := commands.Command[cmd.CalledAs()]
	if permission != "bash" {
		return true, nil
	}

	valid, err := checkToken()
	if err != nil {
		return false, err
	}

	return valid, nil
}

func checkToken() (bool, error) {
	vineAuthToken := path.Join(state.Location.Auth, ".vine_auth_token")
	token, err := utils.ReadLines(vineAuthToken)

	// No token present = error, so just return false here
	if err != nil {
		return false, nil
	}
	if len(token) != 1 || token[0] == "" {
		return false, fmt.Errorf("vine auth token is not set correctly")
	}

	for i := -1; i <= 1; i++ {
		testTime := time.Now().UTC().Add(time.Duration(i) * time.Minute).Format("15:04")
		expectedToken, err := generateAuthToken(testTime)
		if err != nil {
			return false, err
		}

		if token[0] == expectedToken {
			return true, nil
		}
	}

	return false, nil
}
