package proxy

import "net/http"

import (
	"testing"
	//"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/vulcand/oxy/testutils"
	"io/ioutil"
	"net/http/httptest"
)

//func TestTrace(t *testing.T) { TestingT(t) }

type TraceSuite struct{}

//var _ = Suite(&TraceSuite{})

func TestNextHandlerWorks(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	//create some handler
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Length", "5")
		w.Write([]byte("hello"))
	})

	teeHandler, _ := New(nextHandler)

	//trace := &bytes.Buffer{}
	//tee, err := New(teeHandler)

	srv := httptest.NewServer(teeHandler)
	defer srv.Close()

	res, _, err := testutils.MakeRequest(srv.URL+"/db1", testutils.Method("POST"), testutils.Body("123456"))

	if res.StatusCode != http.StatusOK || err != nil {
		t.Fatal("Error posting")
	}

	if res.Header.Get("Content-Length") != "5" {
		t.Error("Content-Length did not pass")
	}

	//buf := &bytes.Buffer{}
	//res.Body.Read(buf)
	//res.Body.Read()
	body, err := ioutil.ReadAll(res.Body)
	log.Debug(body)
}
