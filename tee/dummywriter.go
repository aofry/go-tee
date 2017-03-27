package proxy

import (
	"net/http"
	//log "github.com/Sirupsen/logrus"
)

type DummyResponseWriter struct {
}

func (d *DummyResponseWriter) Header() http.Header {
	someHeader := http.Header{}
	return someHeader
}

func (d *DummyResponseWriter) Write(buf []byte) (int, error) {
	//dummy wrote response
	return len(buf), nil
}

func (d *DummyResponseWriter) WriteHeader(header int) {

}
