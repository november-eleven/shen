package main

import (
	"flag"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/november-eleven/shen/server"
	"github.com/november-eleven/shen/server/container"
	"github.com/november-eleven/shen/server/context"
)

var port uint64

func init() {
	flag.Uint64Var(&port, "port", 3000, "server listening port")
}

func main() {

	flag.Parse()
	if s := os.Getenv("PORT"); s != "" {
		if p, err := strconv.ParseUint(s, 10, 64); err == nil {
			port = p
		}
	}

	exchange := container.NewExchangeContainer()
	peers := container.NewPeersContainer()
	repository := context.NewRepository(exchange, peers)

	err := server.Start(server.Options{
		Exchange:   exchange,
		Peers:      peers,
		Repository: repository,
		Port:       port,
	})

	if err != nil {
		log.Error(err)
		os.Exit(255)
	}

}
