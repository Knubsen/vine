package main

import (
	"github.com/Knubsen/vine/cmd"
	"github.com/Knubsen/vine/models"
	"github.com/Knubsen/vine/state"
)

func main() {
	location := &models.Location{
		Path:    "~/.vine/",
		Scripts: "~/.vine/vine_scripts/",
		Config:  "~/.vine/vine_aliases.yaml",
	}
	state.SetState(location)

	cmd.Execute()
}
