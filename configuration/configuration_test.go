package configuration

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

// setting basic
func TestViper(t *testing.T) {
	var config *viper.Viper = viper.New()
	require.NotNil(t, config)

}

// configuration read file json
func TestConfigJson(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	require.Nil(t, err)

	// test config
	fmt.Println(config.GetString("database.port_db"))
	require.Equal(t, "golang_viper", config.GetString("app.name"))
	require.Equal(t, 3306, config.GetInt("database.port_db"))
	require.Equal(t, "your_db", config.GetString("database.database_db"))

}

// configuration read file .env
func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig()
	require.Nil(t, err)

	// test config
	fmt.Println(config.Password)
	require.Equal(t, "8081", config.Port)
	require.Equal(t, "3306", config.PortDB)
}
