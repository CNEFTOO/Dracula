package app

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/seaung/Dracula/internal/App/store"
	"github.com/seaung/Dracula/pkg/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	recommendHomeDir  = ".Dracula"
	defaultConfigName = "Dracula.yml"
)

var CFG string

func initConfig() {
	if CFG != "" {
		viper.SetConfigFile(CFG)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(filepath.Join(home, recommendHomeDir))
		viper.AddConfigPath(".")
		viper.SetConfigType("yml")
		viper.SetConfigName(defaultConfigName)
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("Dracula")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func initStore() error {
	options := &db.PostgresOptions{
		Host:     viper.GetString("postgres.host"),
		Port:     viper.GetInt("postgres.port"),
		User:     viper.GetString("postgres.user"),
		Password: viper.GetString("postgres.password"),
		DBName:   viper.GetString("postgres.dbname"),
		LogLevel: viper.GetInt("postgres.log_level"),
		MaxIdle:  viper.GetInt("postgres.max_idle"),
		MaxOpen:  viper.GetInt("postgres.max_open"),
		MaxLife:  viper.GetDuration("postgres.max_life"),
	}

	instance, err := db.NewPostgresConnection(options)
	if err != nil {
		return err
	}
	_ = store.NewStore(&instance)

	return nil
}
