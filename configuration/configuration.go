package configuration

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Configuration struct {
	GeoffCdkNotesTableName string
}

func NewConfig() Configuration {
	viperInstance := viper.New()

	viperInstance.SetDefault("GeoffCdkNotesTableName", os.Getenv("GEOFF_CDK_NOTES_TABLE_NAME"))

	var configuration Configuration
	if err := viperInstance.Unmarshal(&configuration); err != nil {
		log.WithField("err", err).Fatal("unable to parse configuration")
	}

	log.WithField("configuration", configuration).Info("configuration supplied")
	return configuration
}
