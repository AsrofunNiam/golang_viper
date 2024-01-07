package configuration

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Port     string `mapstructure:"PORT"`
	PortDB   string `mapstructure:"PORT_DB"`
	Host     string `mapstructure:"HOST_DB"`
	Password string `mapstructure:"PASSWORD_DB"`
	User     string `mapstructure:"USER_DB"`
	Db       string `mapstructure:"DATABASE_DB"`
}

func LoadConfig() (*Configuration, error) {
	var config Configuration

	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".config_env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
