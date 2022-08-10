package configuration

import (
	"api-address-golang/repository/configuration/entities"

	"github.com/spf13/viper"
)

var conf *generalConfiguration

type generalConfiguration struct {
	API entities.ApiConfigurationEntity
	DB  entities.DatabaseConfigurationEntity
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	conf = new(generalConfiguration)

	conf.API = entities.ApiConfigurationEntity{
		Port: viper.GetString("api.port"),
	}

	conf.DB = entities.DatabaseConfigurationEntity{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

func GetDatabase() entities.DatabaseConfigurationEntity {
	return conf.DB
}

func GetServerPort() string {
	return conf.API.Port
}
