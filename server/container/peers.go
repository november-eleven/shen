package container

import (
	"sync"
)

type PeersContainer interface {
	Add(id string)
	Remove(id string)
	Values() []string
}

type SimplePeersContainer struct {
	mutex sync.RWMutex
	db    map[string]bool
}

func NewPeersContainer() PeersContainer {
	return &SimplePeersContainer{db: make(map[string]bool)}
}

func (s *SimplePeersContainer) Add(id string) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.db[id] = true

}

func (s *SimplePeersContainer) Remove(id string) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.db, id)

}

func (s *SimplePeersContainer) Values() []string {

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	l := []string{}

	for id := range s.db {
		l = append(l, id)
	}

	return l
}
