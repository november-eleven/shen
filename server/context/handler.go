package context

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/november-eleven/shen/server/render"
)

type ContextPayload struct {
	ID string `json:"id"`
}

func LoginHandler(repository ContextRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		p := ContextPayload{repository.Login()}
		render.JSON(w, 200, p)
		fmt.Printf("Login: %s\n", p.ID)

	}
}

func LogoutHandler(repository ContextRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			// TODO
			log.Printf("An error has occured: %s", err)
			http.Error(w, "An error has occured", http.StatusInternalServerError)
			return
		}

		p := &ContextPayload{}

		if err := json.Unmarshal(b, p); err != nil {
			// TODO
			log.Printf("An error has occured: %s", err)
			http.Error(w, "An error has occured", http.StatusInternalServerError)
			return
		}

		repository.Logout(p.ID)
		render.NoContent(w)
		fmt.Printf("Logout: %s\n", p.ID)

	}
}
