package render

import (
	"errors"
	"net/http"
	"testing"
)

type responseWriter struct {
	header http.Header
	status int
	buffer []byte
}

func (m *responseWriter) Header() http.Header {
	return m.header
}

func (m *responseWriter) Write(p []byte) (int, error) {
	m.buffer = p
	return len(p), nil
}

func (m *responseWriter) WriteString(s string) (int, error) {
	m.buffer = []byte(s)
	return len(s), nil
}

func (m *responseWriter) WriteHeader(s int) {
	m.status = s
}

func TestNoContent(t *testing.T) {

	w := &responseWriter{header: http.Header{}}

	NoContent(w)

	s := string(w.buffer)

	if w.status != 204 {
		t.Fatalf("unexpected status code: %d", w.status)
	}

	if s != "" {
		t.Fatalf("unexpected body response: %s", s)
	}

}

func TestJSONWithStruct(t *testing.T) {

	w := &responseWriter{header: http.Header{}}

	var fixtures = []struct {
		ID      string `json:"id"`
		Enabled bool   `json:"enabled"`
	}{
		{"1770f3c9-baf8-4f0c-bfa8-79021abc164c", true},
		{"32840865-8f9c-4627-a76e-d9b395446bbe", false},
	}

	JSON(w, 201, fixtures)

	s := string(w.buffer)
	c := w.Header().Get("Content-Type")

	if w.status != 201 {
		t.Fatalf("unexpected status code: %d", w.status)
	}

	if s != `[{"id":"1770f3c9-baf8-4f0c-bfa8-79021abc164c","enabled":true},`+
		`{"id":"32840865-8f9c-4627-a76e-d9b395446bbe","enabled":false}]` {
		t.Fatalf("unexpected body response: %s", s)
	}

	if c != "application/json; charset=utf-8" {
		t.Fatalf("unexpected content type response: %s", c)
	}

}

func TestJSONWithError(t *testing.T) {

	w := &responseWriter{header: http.Header{}}

	err := errors.New("error: bad luck :(")

	JSON(w, 503, err)

	s := string(w.buffer)
	c := w.Header().Get("Content-Type")

	if w.status != 503 {
		t.Fatalf("unexpected status code: %d", w.status)
	}

	if s != `{"message":"error: bad luck :("}` {
		t.Fatalf("unexpected body response: %s", s)
	}

	if c != "application/json; charset=utf-8" {
		t.Fatalf("unexpected content type response: %s", c)
	}

}
