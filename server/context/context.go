package context

import (
	"github.com/november-eleven/shen/server/container"
	"github.com/satori/go.uuid"
)

// Repository define a mecanism to create or delete peers.
type Repository interface {
	Login() string
	Logout(string)
}

// DefaultRepository is a simple and a default implementation of Repository.
type DefaultRepository struct {
	exc container.ExchangeContainer
	pc  container.PeersContainer
}

// Login will create a new peer.
func (r *DefaultRepository) Login() string {
	return uuid.NewV4().String()
}

// Logout will delete and clean a peer.
func (r *DefaultRepository) Logout(id string) {
	r.exc.Flush(id)
	r.pc.Remove(id)
}

// NewRepository will create a new Repository.
func NewRepository(exc container.ExchangeContainer, pc container.PeersContainer) Repository {
	return &DefaultRepository{exc: exc, pc: pc}
}
