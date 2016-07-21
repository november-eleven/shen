package container

import (
	"sync"
)

// PeersContainer define a collection of available peers.
type PeersContainer interface {
	Add(id string)
	Remove(id string)
	Values() []string
}

// SimplePeersContainer is a in-memory PeersContainer.
type SimplePeersContainer struct {
	mutex sync.RWMutex
	db    map[string]bool
}

// NewPeersContainer creates a new PeersContainer
func NewPeersContainer() PeersContainer {
	return &SimplePeersContainer{db: make(map[string]bool)}
}

// Add will append a new available peer.
func (s *SimplePeersContainer) Add(id string) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.db[id] = true

}

// Remove will delete a peer.
func (s *SimplePeersContainer) Remove(id string) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.db, id)

}

// Values return a list of available peers.
func (s *SimplePeersContainer) Values() []string {

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	l := []string{}

	for id := range s.db {
		l = append(l, id)
	}

	return l
}
