package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/aofry/go-tee/proxy"
)

func main() {
	log.SetLevel(log.DebugLevel)

	log.Info("Starting")

	proxy.New()

	log.Info("Ending")
}
