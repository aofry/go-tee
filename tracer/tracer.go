package main

import (
	log "github.com/Sirupsen/logrus"

	"github.com/vulcand/oxy/trace"
	"net/http"
	"os"
)

func main() {
	log.SetLevel(log.DebugLevel)

	log.Info("Starting")

	doNothingHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Length", "5")
		w.Write([]byte("hello"))
	})

	//buf := &bytes.Buffer{}

	tracer, err := trace.New(doNothingHandler, log.StandardLogger().Out)

	if err != nil {
		log.Error("could not init tracer ", err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	s := &http.Server{
		Addr:    (":" + port),
		Handler: tracer,
	}

	log.Info("starting http server on port ", port)

	err = s.ListenAndServe()

	if err != nil {
		log.Error("could not listen: ", err)

	}

	log.Info("Ending")
}
