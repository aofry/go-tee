package proxy

import (
	log "github.com/Sirupsen/logrus"
	"github.com/aofry/go-tee/util"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/utils"
	"io"
	"net/http"
)

type Tee struct {
	errHandler   utils.ErrorHandler
	next         http.Handler
	reqHeaders   []string
	respHeaders  []string
	writer       io.Writer
	requests     chan *http.Request
	debugForward *forward.Forwarder
	debugHost    string
}

type Option func(*Tee) error

func New(next http.Handler) (*Tee, error) {
	//TODO add external param for concurrent limit
	concurrentLimit := 1

	requestsChan := make(chan *http.Request, concurrentLimit)

	//not sending setters so no errors expected
	fw, _ := forward.New()

	t := &Tee{
		next:         next,
		requests:     requestsChan,
		debugForward: fw,
		debugHost:    util.GetenvNoDefault("DEBUG_BACKEND"),
	}

	if t.errHandler == nil {
		t.errHandler = utils.DefaultHandler
	}

	//proxy := http.HandlerFunc(DebugHandler)

	return t, nil
}

func (t *Tee) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	pw := &utils.ProxyWriter{W: w}
	log.Info("Now I'm before the real proxy")
	t.next.ServeHTTP(pw, req)
	log.Info("Now I'm after the real proxy. ", pw.StatusCode(), " ")

	//limit the actual requests that are going out
	if len(t.requests) < cap(t.requests) {
		t.requests <- req
		go t.sendDebugRequest()
	}

}

func (t *Tee) sendDebugRequest() {
	request := <-t.requests

	w := &DummyResponseWriter{}
	//clone request so the original can be free to GC and debug is completly independent
	newRequest := t.copyRequest(request, t.debugHost)
	log.Info(newRequest.Host)
	t.debugForward.ServeHTTP(w, newRequest)
	log.Info("Sent request to debug backend")
}

func (f *Tee) copyRequest(req *http.Request, host string) *http.Request {
	outReq := new(http.Request)
	//*outReq = *req // includes shallow copies of maps, but we handle this below

	outReq.URL = utils.CopyURL(req.URL)
	outReq.URL.Host = host
	outReq.Host = host
	outReq.RequestURI = req.RequestURI
	//outReq.URL.Opaque = req.RequestURI
	// raw query is already included in RequestURI, so ignore it to avoid dupes
	//outReq.URL.RawQuery = ""

	outReq.Proto = "HTTP/1.1"
	outReq.ProtoMajor = 1
	outReq.ProtoMinor = 1

	// Overwrite close flag so we can keep persistent connection for the backend servers
	outReq.Close = false

	outReq.Header = make(http.Header)
	utils.CopyHeaders(outReq.Header, req.Header)

	return outReq
}
