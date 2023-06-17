package config

import (
	"fmt"

	"github.com/Epritka/morpheus/config"

	"github.com/Epritka/morpheus/ogm"
)

func newDatabseConfig() *config.Config {
	return &config.Config{
		URI:      settings.Database.Uri,
		Username: settings.Database.User,
		Password: settings.Database.Password,
	}
}

var conn *ogm.Driver

func newDatabaseConnection(config *config.Config) *ogm.Driver {
	database, err := ogm.Connect(config)
	if err != nil {
		panic(fmt.Errorf("fatal error connect database: %w", err))
	}
	return database
}

func InitDatabase() {
	config := newDatabseConfig()
	conn = newDatabaseConnection(config)
}

func GetDatabaseConnection() *ogm.Driver {
	return conn
}
