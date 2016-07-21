package peers

import (
	"encoding/json"
	"errors"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/november-eleven/shen/server/container"
	"github.com/november-eleven/shen/server/render"
	"github.com/pressly/chi"
)

var (
	// ErrExchangeRead is returned when an error has occurred while reading request body.
	ErrExchangeRead = errors.New("Cannot read request as Peer Exchange")

	// ErrExchangeParse is returned when an error has occurred while parsing json payload.
	ErrExchangeParse = errors.New("Cannot read request as Peer Exchange")
)

// RegisterHandler will push a peer as alive and ready to be linked with other peers.
func RegisterHandler(c container.PeersContainer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.WithField("peers", "register").Error(err)
			render.JSON(w, http.StatusBadRequest, ErrExchangeRead)
			return
		}

		id := string(b)

		c.Add(id)
		render.NoContent(w)
		log.WithField("peers", "register").Info(id)

	}
}

// ListHandler will display a list of available peers.
func ListHandler(c container.PeersContainer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, 200, c.Values())
	}
}

// PullHandler will display a list of available peers which request a handshake.
func PullHandler(c container.ExchangeContainer) func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(ctx, "id")
		render.JSON(w, 200, c.Values(id))
		log.WithField("peers", "pull").Info(id)

	}
}

// PushHandler will push a handshake request.
func PushHandler(c container.ExchangeContainer) func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(ctx, "id")

		p, err := parse(r)
		if err != nil {
			render.JSON(w, http.StatusBadRequest, err)
			return
		}

		c.Add(id, *p)
		render.NoContent(w)
		log.WithField("peers", "push").Info(id)

	}
}

// RemoveHandler will delete an awaiting handshake.
func RemoveHandler(c container.ExchangeContainer) func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(ctx, "id")

		p, err := parse(r)
		if err != nil {
			render.JSON(w, http.StatusBadRequest, err)
			return
		}

		c.Remove(id, *p)
		render.NoContent(w)
		log.WithField("peers", "remove").Info(id)

	}
}

func parse(r *http.Request) (*container.ExchangePayload, error) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithField("peers", "parse").Error(err)
		return nil, ErrExchangeRead
	}

	p := &container.ExchangePayload{}

	err = json.Unmarshal(b, p)
	if err != nil {
		log.WithField("peers", "parse").Error(err)
		return nil, ErrExchangeParse
	}

	return p, nil

}
