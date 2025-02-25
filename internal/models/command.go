package models

type Command struct {
	Aliases []string `yaml:"aliases"`
	Command string   `yaml:"command"`
}
