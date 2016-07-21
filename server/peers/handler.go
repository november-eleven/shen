package peers

import (
	"encoding/json"
	"fmt"
	"github.com/pressly/chi"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/november-eleven/shen/server/container"
	"github.com/november-eleven/shen/server/render"
)

func RegisterHandler(c container.PeersContainer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// TODO
			log.Printf("An error has occured: %s", err)
			http.Error(w, "An error has occured", http.StatusInternalServerError)
			return
		}

		id := string(b)

		c.Add(id)
		fmt.Printf("Register: %s\n", id)

		render.NoContent(w)

	}
}

func ListHandler(c container.PeersContainer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, 200, c.Values())
	}
}

func PullHandler(c container.ExchangeContainer) func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(ctx, "id")
		render.JSON(w, 200, c.Values(id))
		fmt.Printf("Pull: %s\n", id)

	}
}

func PushHandler(c container.ExchangeContainer) func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(ctx, "id")

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// TODO
			log.Printf("An error has occured: %s", err)
			http.Error(w, "An error has occured", http.StatusInternalServerError)
			return
		}

		p := container.ExchangePayload{}

		err = json.Unmarshal(b, &p)
		if err != nil {
			// TODO
			log.Printf("An error has occured: %s", err)
			http.Error(w, "An error has occured", http.StatusInternalServerError)
			return
		}

		c.Add(id, p)
		render.NoContent(w)
		fmt.Printf("Push: %s\n", id)

	}
}

func RemoveHandler(c container.ExchangeContainer) func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(ctx, "id")

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// TODO
			log.Printf("An error has occured: %s", err)
			http.Error(w, "An error has occured", http.StatusInternalServerError)
			return
		}

		p := container.ExchangePayload{}

		err = json.Unmarshal(b, &p)
		if err != nil {
			// TODO
			log.Printf("An error has occured: %s", err)
			http.Error(w, "An error has occured", http.StatusInternalServerError)
			return
		}

		c.Remove(id, p)
		render.NoContent(w)
		fmt.Printf("Remove: %s\n", id)

	}
}
