package proxy

import "net/http"

import (
	"testing"
	//"bytes"
	"github.com/vulcand/oxy/testutils"
	"net/http/httptest"
)

//func TestTrace(t *testing.T) { TestingT(t) }

type TraceSuite struct{}

//var _ = Suite(&TraceSuite{})

func TestNextHandlerWorks(t *testing.T) {
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

	res, _, err := testutils.MakeRequest(srv.URL+"/hello", testutils.Method("POST"), testutils.Body("123456"))

	if res.StatusCode != http.StatusOK || err != nil {
		t.Fatal("Error posting")
	}
}
