package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/aofry/go-tee/proxy"
)

func main() {
	log.SetLevel(log.InfoLevel)

	log.Println("Starting")

	proxy.New()

	log.Println("Ending")
}
