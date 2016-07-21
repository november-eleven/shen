package context

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/november-eleven/shen/server/render"
)

var (
	// ErrPeerRead is returned when an error has occurred while reading request body.
	ErrPeerRead = errors.New("Cannot read request as Peer")

	// ErrPeerParse is returned when an error has occurred while parsing json payload.
	ErrPeerParse = errors.New("Cannot parse request as Peer")
)

// Peer define a remote client.
type Peer struct {
	ID string `json:"id"`
}

// LoginHandler will handle peers creation.
func LoginHandler(repository Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		p := Peer{repository.Login()}
		render.JSON(w, 200, p)
		log.WithField("context", "login").Info(p.ID)

	}
}

// LogoutHandler will handle peers deletion.
func LogoutHandler(repository Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.WithField("context", "logout").Error(err)
			render.JSON(w, http.StatusBadRequest, ErrPeerRead)
			return
		}

		p := &Peer{}

		if err := json.Unmarshal(b, p); err != nil {
			log.WithField("context", "logout").Error(err)
			render.JSON(w, http.StatusBadRequest, ErrPeerParse)
			return
		}

		repository.Logout(p.ID)
		render.NoContent(w)
		log.WithField("context", "logout").Info(p.ID)

	}
}
