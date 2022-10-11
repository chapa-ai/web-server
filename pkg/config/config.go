package config

import (
	"github.com/spf13/viper"
	"sync"
)

var once sync.Once

type Config struct {
	DBDriver    string `mapstructure:"DB_DRIVER"`
	DBSource    string `mapstructure:"DB_SOURCE"`
	ConfigToken string `json:"configToken"`
	Port        string `json:"port"`
	Once        sync.Once
}

type DbConnection struct{}

var (
	dbConnOnce sync.Once
	conn       *Config
)

func GetConfig() *Config {
	dbConnOnce.Do(func() {
		conn = &Config{}
		err := conn.LoadConfig(".")
		if err != nil {
			panic(err)
		}
	})
	return conn
}

func (c *Config) LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config/config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(c)
	if err != nil {
		return err
	}
	return err
}
