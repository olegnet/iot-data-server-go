package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const DefaultConfigName = "config.json"

type Config struct {
	Bind     string         `json:"bind"`
	Postgres PostgresConfig `json:"postgres"`
}

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
	Dbname   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (config *Config) String() string {
	marshal, err := json.Marshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	return string(marshal)
}

func (config *Config) PostgresConfig() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.Password,
		config.Postgres.Dbname)
}

func Read(configName string) Config {
	var config Config

	file, errOpen := os.Open(configName)
	if errOpen != nil {
		log.Fatal(errOpen)
	}

	errDecode := json.NewDecoder(file).Decode(&config)
	if errDecode != nil {
		log.Fatal(errDecode)
	}

	_ = file.Close()

	log.Printf("config=%v", config)

	return config
}
