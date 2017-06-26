package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/Sirupsen/logrus"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "book downloader",
	Short: "download book from server",
}

var configFilePath string

func init() {
	// Register a global option for config file path
	RootCmd.PersistentFlags().StringVarP(
		&configFilePath, "config", "c", "", "Path to the config file",
	)
}

func LoadConfiguration(cmd *cobra.Command, args []string) {
	// Load the configuration
	if configFilePath != "" {
		viper.SetConfigFile(configFilePath)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	}

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		if viper.ConfigFileUsed() == "" {
			log.Fatalf("Unable to find configuration file.")
		}

		log.Fatalf("Failed to load %s: %v", viper.ConfigFileUsed(), err)
	} else {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}
