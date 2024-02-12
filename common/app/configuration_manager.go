package app

import "country_information_api/common/postgre"

type ConfigurationManager struct {
	PostgreConfig postgre.Config
}

func NewConfigurationManager() *ConfigurationManager {
	postgreConfig := getPostgreConfig()
	return &ConfigurationManager{
		PostgreConfig: postgreConfig,
	}
}

func getPostgreConfig() postgre.Config {
	return postgre.Config{
		Host:                  "localhost",
		Port:                  "7432",
		DbName:                "countryinfoapp",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	}
}
