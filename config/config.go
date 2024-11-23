package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return errors.Join(
		os.Setenv("APP_NAME", viper.GetString("app-name")),
		os.Setenv("APP_VERSION", viper.GetString("app-version")),
		os.Setenv("PORT", viper.GetString("port")),
		os.Setenv("BASE_URL", viper.GetString("base-url")),
	)
}
