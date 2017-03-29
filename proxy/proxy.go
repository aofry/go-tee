package proxy

import (
	goTee "github.com/aofry/go-tee/tee"
	"github.com/aofry/go-tee/util"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"
	"net/http"
	"os"
	//	log "github.com/Sirupsen/logrus"
)

func New() {
	proxyPort := util.Getenv("PORT", "8080")

	proxy := http.HandlerFunc(ProxyHandler)
	teeHandler, _ := goTee.New(proxy)

	// that's it! our reverse proxy is ready!
	s := &http.Server{
		Addr:    (":" + proxyPort),
		Handler: teeHandler,
	}
	s.ListenAndServe()
}

func ProxyHandler(w http.ResponseWriter, req *http.Request) {
	service := os.Getenv("REAL_BACKEND")

	// let us forward this request to another server
	req.URL = testutils.ParseURI(service)

	// Forwards incoming requests to whatever location URL points to, adds proper forwarding headers
	fwd, _ := forward.New()

	fwd.ServeHTTP(w, req)
}

//func TeeHandler(w http.ResponseWriter, req *http.Request) {
//
//}
