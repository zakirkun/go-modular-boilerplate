package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	filename string
}

func NewConfig(filename string) Config {
	return Config{filename: filename}
}
func (c *Config) Initialize() error {

	configName := filepath.Base(c.filename)

	configExtension := filepath.Ext(c.filename)
	configExtension = strings.TrimPrefix(configExtension, ".")

	viper.SetConfigName(configName)
	viper.SetConfigType(configExtension)
	viper.AddConfigPath(filepath.Dir(c.filename))

	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return err
		}
		return err
	}

	return nil
}

func checkKey(key string) {
	if !viper.IsSet(key) {
		log.Fatalf("Configuration key %s not found; aborting \n", key)
		os.Exit(1)
	}
}

func GetString(key string) string {
	checkKey(key)
	return viper.GetString(key)
}

func GetInt(key string) int {
	checkKey(key)
	return viper.GetInt(key)
}

func GetBool(key string) bool {
	checkKey(key)
	return viper.GetBool(key)
}
