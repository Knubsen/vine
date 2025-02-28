package models

type AuthorizedCommand struct {
	Command map[string]string `json:"commands"`
}

// type CommandAuthConfig struct {
// 	AuthCommands []AuthorizedCommand `json:"authconfig"`
// }
