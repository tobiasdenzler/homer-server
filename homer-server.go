package main

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"github.com/tobiasdenzler/homer-server/config"
	"github.com/tobiasdenzler/homer-server/pkg/dss"
)

func init() {
	loglevel, _ := logrus.ParseLevel(config.Config.Log.Level)
	log.SetLevel(loglevel)
}

func main() {

	log.Info("Starting application")

	dss.Call("/json/property/query", map[string]string{"query": "/apartment/zones/*(*)"})

}
