package proxy

import (
	"os"
	"fmt"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"
	"net/http"
	"time"
)

func New() {
	proxyPort := os.Getenv("PORT")

	if proxyPort == "" {
		proxyPort = "8080"
	}

	redirect := http.HandlerFunc(ProxyHandler)

	// that's it! our reverse proxy is ready!
	s := &http.Server{
		Addr:    (":" + proxyPort),
		Handler: redirect,
	}
	s.ListenAndServe()
}

func ProxyHandler (w http.ResponseWriter, req *http.Request) {
	service := os.Getenv("SERVICE")

	now := time.Now()
	fmt.Println(now, "Got a request ")
	// let us forward this request to another server
	req.URL = testutils.ParseURI(service)

	// Forwards incoming requests to whatever location URL points to, adds proper forwarding headers
	fwd, _ := forward.New()

	fwd.ServeHTTP(w, req)
}