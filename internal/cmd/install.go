package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Knubsen/vine/internal/state"
	"github.com/Knubsen/vine/internal/utils"
	"github.com/spf13/cobra"
)

const startDelimiter = "# Vine stuff"
const endDelimiter = "# ffuts eniV"

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "installs vine",
	Long:  "creates the necessary directories and adds the vine function to the .bashrc",
	// Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lines, err := utils.ReadLines(state.Location.BashRC)
		if err != nil {
			fmt.Printf("Error reading the lines: %v", err)
			return
		}

		if checkForDelimiter(lines) {
			fmt.Printf(".bashrc already contains vine\n")
			return
		}

		err = addVine()
		if err != nil {
			fmt.Printf("Error adding vine to .bashrc: %v", err)
			return
		}
		fmt.Println("Vine added to .bashrc")
	},
}

func addVine() error {
	file, err := os.OpenFile(state.Location.BashRC, os.O_APPEND|os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	bashrcTemplate, err := getBashrcTemplate()
	if err != nil {
		return err
	}

	_, err = file.WriteString(bashrcTemplate)
	if err != nil {
		return fmt.Errorf("error writing to file: %e", err)
	}

	return nil
}

func getBashrcTemplate() (string, error) {
	var outputString = "\n"

	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return "", fmt.Errorf("Error getting working directory: %e", err)
	}

	file, err := utils.ReadLines(filepath.Join(rootDir, "config", ".bashrc"))
	if err != nil {
		return "", fmt.Errorf("Error reading .bashrc template: %e", err)
	}

	for _, line := range file {
		outputString += line + "\n"
	}

	return outputString, nil
}

func checkForDelimiter(lines []string) bool {
	for _, line := range lines {
		if strings.TrimSpace(line) == strings.TrimSpace(startDelimiter) {
			return true
		}
	}

	return false
}

func init() {
	rootCmd.AddCommand(installCmd)
}
