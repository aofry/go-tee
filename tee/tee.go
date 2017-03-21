package proxy

import (
	log "github.com/Sirupsen/logrus"
	"github.com/vulcand/oxy/utils"
	"io"
	"net/http"
	"github.com/vulcand/oxy/forward"
	"net/url"
)

type Tee struct {
	errHandler  utils.ErrorHandler
	next        http.Handler
	reqHeaders  []string
	respHeaders []string
	writer      io.Writer
	requests    chan *http.Request
	debugForward *forward.Forwarder
	debugUrl    string
}

type Option func(*Tee) error

func New(next http.Handler, opts ...Option) (*Tee, error) {
	requestsChan := make(chan *http.Request, 100)
	fw, _ := forward.New()
	//TODO add url for debug backend
	t := &Tee{
		next:     next,
		requests: requestsChan,
		debugForward: fw,
	}
	for _, o := range opts {
		if err := o(t); err != nil {
			return nil, err
		}
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
	t.requests <- req

	go t.sendDebugRequest()
}

func (t *Tee) sendDebugRequest() {
	//request := <-t.requests

	//w := DummyResponseWriter {}
	//var pUrl *url.URL = &url.URL{}
	//pUrl, _ = pUrl.Parse(t.debugUrl)
	//newRequest := t.copyRequest(request, pUrl)
	//t.debugForward.ServeHTTP(w, newRequest)

}

func (f *Tee) copyRequest(req *http.Request, u *url.URL) *http.Request {
	outReq := new(http.Request)
	*outReq = *req // includes shallow copies of maps, but we handle this below

	outReq.URL = utils.CopyURL(req.URL)
	outReq.URL.Scheme = u.Scheme
	outReq.URL.Host = u.Host
	outReq.URL.Opaque = req.RequestURI
	// raw query is already included in RequestURI, so ignore it to avoid dupes
	outReq.URL.RawQuery = ""

	outReq.Proto = "HTTP/1.1"
	outReq.ProtoMajor = 1
	outReq.ProtoMinor = 1

	// Overwrite close flag so we can keep persistent connection for the backend servers
	outReq.Close = false

	outReq.Header = make(http.Header)
	utils.CopyHeaders(outReq.Header, req.Header)

	return outReq
}