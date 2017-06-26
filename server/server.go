package server

import (
	"druid/black/backman/handlers"

	"github.com/spf13/viper"

//	"time"

	log "github.com/Sirupsen/logrus"
)

// Server is main loop service
func Server() {
	cmd := viper.GetString("cmd")


	log.Info("Start cmd:", cmd)
	mapMux[cmd]()
}
