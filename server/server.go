package main

import (
	"net/http"

	"github.com/november-eleven/shen/server/container"
	"github.com/november-eleven/shen/server/context"
	"github.com/november-eleven/shen/server/peers"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

type Options struct {
	peers      container.PeersContainer
	exchange   container.ExchangeContainer
	repository context.ContextRepository
}

func Start(o Options) {

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CloseNotify)

	r.Route("/api/peers", func(r chi.Router) {

		r.Get("/", peers.ListHandler(o.peers))
		r.Post("/", peers.RegisterHandler(o.peers))

		r.Route("/:id", func(r chi.Router) {

			r.Get("/", peers.PullHandler(o.exchange))
			r.Post("/", peers.PushHandler(o.exchange))
			r.Delete("/", peers.RemoveHandler(o.exchange))

		})

	})

	r.Post("/api/login", context.LoginHandler(o.repository))
	r.Post("/api/logout", context.LogoutHandler(o.repository))

	r.Get("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("."))))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.ListenAndServe(":3000", r)

}

type Foo struct {
}

func (f Foo) Server() {

}
