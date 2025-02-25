package parser

import (
	"io"
	"os"

	"github.com/Knubsen/vine/internal/state"
)

func loadAliases() ([]byte, error) {
	file, err := os.Open(state.Location.Config)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}
