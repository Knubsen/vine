package main

import (
	"github.com/Knubsen/vine/internal/cmd"
	"github.com/Knubsen/vine/internal/models"
	"github.com/Knubsen/vine/internal/state"
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
