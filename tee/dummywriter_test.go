package proxy

import (
	"net/http"
	"testing"
)

func TestWriterHasHeaderStruct(t *testing.T) {
	w := &DummyResponseWriter{}
	if w.Header() == nil {
		t.Error("Could not get writer headers")
	}
}

func TestWrite(t *testing.T) {
	w := &DummyResponseWriter{}
	buf := make([]byte, 100)
	resSize, err := w.Write(buf)

	if err != nil {
		t.Error("Got error writing", err)
	}
	if resSize != len(buf) {
		t.Error("Could not write into dummywriter")
	}
}

func TestWriteError(t *testing.T) {
	w := &DummyResponseWriter{}

	w.WriteHeader(http.StatusOK)
}
