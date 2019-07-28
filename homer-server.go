package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"github.com/tobiasdenzler/homer-server/config"
	"github.com/tobiasdenzler/homer-server/pkg/dss"
)

func init() {
	// set loglevel based on config
	loglevel, _ := logrus.ParseLevel(config.Config.Log.Level)
	log.SetLevel(loglevel)
}

func main() {

	log.Info("Starting application")

	//dss.Call("/json/property/query", map[string]string{"query": "/apartment/zones/*(*)"})

	// This will start polling and handling events from DSS asynchronously
	var eventChannel = make(chan []byte)
	go dss.PollEvents(42, eventChannel)
	go dss.HandleEvents(eventChannel)

	// keep alive
	var input string
	fmt.Scanln(&input)
}
