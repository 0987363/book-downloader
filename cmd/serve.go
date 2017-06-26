package cmd

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:    "serve",
	Short:  "Start book downloader",
	PreRun: LoadConfiguration,
	Run:    serve,
}

func init() {
	RootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringP(
		"directory",
		"d",
		"",
		"xiaoshuo directory url",
	)
	viper.BindPFlag("directory", serveCmd.Flags().Lookup("directory"))
}

func serve(cmd *cobra.Command, args []string) {
	// Try to connect to the database
	if err := middleware.ConnectDB(
		viper.GetString("database"),
	); err != nil {
		log.Fatalln("connect to db: ", err)
	}

	server.Server()
}

