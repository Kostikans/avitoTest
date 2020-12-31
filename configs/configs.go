package configs

import (
	"os"

	"github.com/spf13/viper"
)

type postgresConfig struct {
	User     string
	Password string
	DBName   string
}

var BdConfig postgresConfig

func Init() {
	BdConfig = postgresConfig{
		User:     os.Getenv("PostgresAvitoUser"),
		Password: os.Getenv("PostgresAvitoPassword"),
		DBName:   os.Getenv("PostgresAvitoDBName"),
	}
}

func ExportConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
