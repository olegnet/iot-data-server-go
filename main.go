package main

import (
	"config"
	"network"
	"postgres"
)

func main() {
	appConfig := config.Read(config.DefaultConfigName)

	database := postgres.Open(appConfig.PostgresConfig())

	network.Start(appConfig.Bind, database)
}
