package render

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

// ErrorWrap wrap an error to expose it as a JSON.
type ErrorWrap struct {
	Err string `json:"message"`
}

// Wrap will wrap the given error as ErrorWrap.
func Wrap(err error) ErrorWrap {
	return ErrorWrap{err.Error()}
}

// JSON will render a JSON on the given ResponseWriter.
func JSON(w http.ResponseWriter, status int, v interface{}) {

	if p, ok := v.(error); ok {
		v = Wrap(p)
	}

	b, err := json.Marshal(v)
	if err != nil {

		log.WithField("render", "json").Error(err)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)

		if b, err = json.Marshal(Wrap(err)); err == nil {
			if _, err := w.Write(b); err != nil {
				log.WithField("render", "write").Error(err)
			}
		}

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if _, err := w.Write(b); err != nil {
		log.WithField("render", "write").Error(err)
	}
}

// NoContent will render nothing.
func NoContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNoContent)
}
