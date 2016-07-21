package main

import (
	"github.com/november-eleven/shen/server/container"
	"github.com/november-eleven/shen/server/context"
)

func main() {

	exchange := container.NewExchangeContainer()
	peers := container.NewPeersContainer()
	repository := context.NewContextRepository(exchange, peers)

	Start(Options{
		exchange:   exchange,
		peers:      peers,
		repository: repository,
	})

}
