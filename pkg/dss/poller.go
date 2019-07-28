package dss

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/tobiasdenzler/homer-server/config"
)

// PollEvents will poll for the subscribed events.
// If we get an event we will notify the events channel.
func PollEvents(subscriptionID int, eventChannel chan []byte) {

	// Subscribe to callScene events sent by the DSS
	Call("/json/event/subscribe", map[string]string{"name": "callScene", "subscriptionID": strconv.Itoa(subscriptionID)})

	// Start polling forever
	for {
		log.Trace("Polling..")

		// Here we call the API with e defines timeout. We will receive any events happening before the timeout.
		data := Call("/json/event/get", map[string]string{"subscriptionID": strconv.Itoa(subscriptionID), "timeout": strconv.Itoa(config.Config.Server.WaitPolling)})

		// Send the JSON event data on the channel
		eventChannel <- data
	}
}
