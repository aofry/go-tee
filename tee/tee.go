package proxy

import (
	log "github.com/Sirupsen/logrus"
	"github.com/vulcand/oxy/utils"
	"io"
	"net/http"
)

type Tee struct {
	errHandler  utils.ErrorHandler
	next        http.Handler
	reqHeaders  []string
	respHeaders []string
	writer      io.Writer
	requests    chan *http.Request
}

type Option func(*Tee) error

func New(next http.Handler, opts ...Option) (*Tee, error) {
	requestsChan := make(chan *http.Request, 100)
	t := &Tee{
		next:     next,
		requests: requestsChan,
	}
	for _, o := range opts {
		if err := o(t); err != nil {
			return nil, err
		}
	}
	if t.errHandler == nil {
		t.errHandler = utils.DefaultHandler
	}
	return t, nil
}

func (t *Tee) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	pw := &utils.ProxyWriter{W: w}
	log.Info("Now I'm before the real proxy")
	t.next.ServeHTTP(pw, req)
	log.Info("Now I'm after the real proxy. ", pw.StatusCode(), " ")
	//t.requests <- req

}
