package dss

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tobiasdenzler/homer-server/config"
)

// HandleEvents will receive DSS events through a channel.
func HandleEvents(eventChannel chan []byte) {
	for {
		event := <-eventChannel
		log.Tracef("Received channel event -> %s", event)
		time.Sleep(time.Duration(config.Config.Server.WaitPolling) * time.Second)
	}
}
