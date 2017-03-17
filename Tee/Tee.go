package proxy

import (
	"net/http"
	"github.com/vulcand/oxy/utils"
	"io"
	"time"
)

type Tee struct {
	errHandler  utils.ErrorHandler
	next        http.Handler
	reqHeaders  []string
	respHeaders []string
	writer      io.Writer
}

type Option func(*Tee) error

func New(next http.Handler, writer io.Writer, opts ...Option) (*Tee, error) {
	t := &Tee{
		writer: writer,
		next:   next,
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
	t.next.ServeHTTP(pw, req)
}