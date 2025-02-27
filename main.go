package main

import (
	"github.com/Knubsen/vine/cmd"
	"github.com/Knubsen/vine/internal/models"
	"github.com/Knubsen/vine/internal/state"
)

func main() {
	location := &models.Location{
		Path:    "~/.vine/",
		BashRC:  "~/.bashrc",
		Scripts: "~/.vine/vine_scripts/",
		Config:  "~/.vine/vine_aliases.yaml",
	}
	state.SetState(location)

	cmd.Execute()
}
