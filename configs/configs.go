package configs

import (
	"github.com/spf13/viper"
	"os"
)


type postgresConfig struct {
	User     string
	Password string
	DBName   string
}

var BdConfig postgresConfig

func Init() {
	BdConfig = postgresConfig{
		User:     os.Getenv("PostgresUser"),
		Password: os.Getenv("PostgresPassword"),
		DBName:   os.Getenv("PostgresDBName"),
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
