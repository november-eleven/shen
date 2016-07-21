package context

import (
	"github.com/november-eleven/shen/server/container"
	"github.com/satori/go.uuid"
)

type ContextRepository interface {
	Login() string
	Logout(string)
}

type DefaultContextRepository struct {
	exc container.ExchangeContainer
	pc  container.PeersContainer
}

func (r *DefaultContextRepository) Login() string {
	return uuid.NewV4().String()
}

func (r *DefaultContextRepository) Logout(id string) {
	r.exc.Flush(id)
	r.pc.Remove(id)
}

func NewContextRepository(exc container.ExchangeContainer, pc container.PeersContainer) ContextRepository {
	return &DefaultContextRepository{exc: exc, pc: pc}
}
