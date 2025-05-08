package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func InitConfigEnvironment() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	//Set default configurations
	viper.SetDefault("app.server.host", "127.0.0.1")
	viper.SetDefault("app.server.port", 4200)
	viper.SetDefault("app.debug.enable", false)

	// Read configuration from file or environment variables
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

}
