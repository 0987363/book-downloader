package main

import (
	log "github.com/Sirupsen/logrus"

	"runtime"
	"time"
)

// Exported onstants for storing build information
var (
	BuildVersion string
	BuildDate    string
	BuildCommit  string
)

func init() {
	time.Local = time.UTC
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	cmd.BuildInfo.Version = BuildVersion
	cmd.BuildInfo.Date = BuildDate
	cmd.BuildInfo.Commit = BuildCommit

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}



}
