package parser

import (
	"fmt"

	"github.com/Knubsen/vine/models"
	"gopkg.in/yaml.v3"
)

func GetConfig() (*models.Config, error) {
	var config models.Config

	yamlContent, err := loadAliases()
	if err != nil {
		return nil, fmt.Errorf("Error: %e", err)
	}

	err = yaml.Unmarshal([]byte(yamlContent), &config)
	if err != nil {
		return nil, fmt.Errorf("Error: %e", err)
	}

	return &config, nil
}
