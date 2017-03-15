package proxy

import (
	"fmt"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"
	"net/http"
	"time"
)

func New() {
	// Forwards incoming requests to whatever location URL points to, adds proper forwarding headers
	fwd, _ := forward.New()

	redirect := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		now := time.Now()
		fmt.Println(now, "Got a request ")
		// let us forward this request to another server
		req.URL = testutils.ParseURI("http://localhost:63450")
		fwd.ServeHTTP(w, req)
	})

	// that's it! our reverse proxy is ready!
	s := &http.Server{
		Addr:    ":8080",
		Handler: redirect,
	}
	s.ListenAndServe()
}
