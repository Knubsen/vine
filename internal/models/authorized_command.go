package models

type AuthorizedCommand struct {
	Command map[string]string `json:"command"`
}
type CommandAuthConfig struct {
	AuthCommands []AuthorizedCommand `json:"authconfig"`
}
