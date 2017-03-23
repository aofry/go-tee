package proxy

import "net/http"

type DummyResponseWriter struct {
}

func (d *DummyResponseWriter) Header() http.Header {
	someHeader := http.Header{}
	return someHeader
}

func (d *DummyResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (d *DummyResponseWriter) WriteHeader(int) {

}
