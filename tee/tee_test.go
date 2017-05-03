package proxy

import (
	log "github.com/Sirupsen/logrus"
	"github.com/vulcand/oxy/testutils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

//func TestTrace(t *testing.T) { TestingT(t) }

type TeeSuite struct {
	generalHandler http.HandlerFunc
	server         *httptest.Server
}

//var _ = Suite(&TeeSuite{})
func TestMain(m *testing.M) {
	log.SetLevel(log.DebugLevel)

	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestNextHandlerWorks(t *testing.T) {
	bodyToSend := "123456"
	os.Setenv("DEBUG_BACKEND", "debug:80")
	dummyHandlerResponseText := "hello"

	//create some handler
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Length", strconv.Itoa(len(dummyHandlerResponseText)))

		w.Write([]byte(dummyHandlerResponseText))
	})

	teeHandler, _ := New(nextHandler)

	//trace := &bytes.Buffer{}
	//tee, err := New(teeHandler)

	srv := httptest.NewServer(teeHandler)
	defer srv.Close()

	res, _, err := testutils.MakeRequest(srv.URL+"/db1", testutils.Method("POST"), testutils.Body(bodyToSend))

	if res.StatusCode != http.StatusOK || err != nil {
		log.Error("posting to tee handler returned: ", res.StatusCode, err)
		t.Fatal("Error posting")
	}

	if res.Header.Get("Content-Length") != strconv.Itoa(len(dummyHandlerResponseText)) {
		t.Error("Content-Length did not pass")
	}

	body, err := ioutil.ReadAll(res.Body)
	log.Debug("body", body)
}

//TODO add test for copy request and that it contain another request object with validation to some fields
