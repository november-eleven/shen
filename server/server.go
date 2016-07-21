package server

import (
	"fmt"
	ctx "golang.org/x/net/context"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/november-eleven/shen/server/container"
	"github.com/november-eleven/shen/server/context"
	"github.com/november-eleven/shen/server/peers"
	"github.com/november-eleven/shen/server/render"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

// Options define service dependency.
type Options struct {
	Peers      container.PeersContainer
	Exchange   container.ExchangeContainer
	Repository context.Repository
	Port       uint64
}

// Start will launch the shen service.
func Start(o Options) error {

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CloseNotify)

	r.Route("/api/peers", func(r chi.Router) {

		r.Get("/", peers.ListHandler(o.Peers))
		r.Post("/", peers.RegisterHandler(o.Peers))

		r.Route("/:id", func(r chi.Router) {

			r.Get("/", peers.PullHandler(o.Exchange))
			r.Post("/", peers.PushHandler(o.Exchange))
			r.Delete("/", peers.RemoveHandler(o.Exchange))

		})

	})

	r.Post("/api/login", context.LoginHandler(o.Repository))
	r.Post("/api/logout", context.LogoutHandler(o.Repository))

	r.FileServer("/assets/", http.Dir("."))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	r.NotFound(func(ctx ctx.Context, w http.ResponseWriter, r *http.Request) {
		render.JSON(w, http.StatusNotFound, fmt.Errorf("page not found"))
	})

	addr := fmt.Sprintf(":%d", o.Port)

	log.Infof("Listening on %s\n", addr)
	return http.ListenAndServe(addr, r)

}
