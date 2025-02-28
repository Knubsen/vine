package state

import (
	"log"
	"os"
	"strings"

	"github.com/Knubsen/vine/internal/models"
)

var Location *models.Location
var HomeDir string

func SetState(location *models.Location) {
	if HomeDir == "" {
		setHomeDir()
	}
	Location = &models.Location{
		Path:    strings.ReplaceAll(location.Path, "~", HomeDir),
		BashRC:  strings.ReplaceAll(location.BashRC, "~", HomeDir),
		Config:  strings.ReplaceAll(location.Config, "~", HomeDir),
		Auth:    strings.ReplaceAll(location.Auth, "~", HomeDir),
		Scripts: strings.ReplaceAll(location.Scripts, "~", HomeDir),
	}
}

func setHomeDir() {
	var err error
	HomeDir, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
}
