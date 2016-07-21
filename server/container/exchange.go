package container

import (
	"encoding/json"
	"sync"
)

// ExchangePayloadLimit define the maximum of waiting peers.
const ExchangePayloadLimit = 256

// ExchangePayload define a peer waiting to connect.
type ExchangePayload struct {
	ID  string          `json:"id"`
	Raw json.RawMessage `json:"signal"`
}

// ExchangeContainer handle ExchangePayload storage.
type ExchangeContainer interface {
	Add(id string, payload ExchangePayload)
	Remove(id string, payload ExchangePayload)
	Flush(id string)
	Values(id string) []ExchangePayload
}

// SimpleExchangeContainer is a in-memory ExchangeContainer.
// It uses a Read/Write mutex and split peer into several containers for concurrent access.
type SimpleExchangeContainer struct {
	mutex sync.RWMutex
	db    map[string](*exchangePayloadContainer)
}

// NewExchangeContainer creates a new ExchangeContainer.
func NewExchangeContainer() ExchangeContainer {
	return &SimpleExchangeContainer{db: make(map[string](*exchangePayloadContainer))}

}

// Add will append a new ExchangePayload for a peer.
func (s *SimpleExchangeContainer) Add(id string, e ExchangePayload) {

	s.mutex.Lock()

	l, ok := s.db[id]

	if !ok {
		l = &exchangePayloadContainer{db: make(map[string]ExchangePayload)}
		s.db[id] = l
	}

	s.mutex.Unlock()

	l.Add(e)

}

// Remove will delete a ExchangePayload for a peer.
func (s *SimpleExchangeContainer) Remove(id string, e ExchangePayload) {

	s.mutex.Lock()
	l, ok := s.db[id]
	s.mutex.Unlock()

	if ok {
		l.Remove(e)
	}

}

// Flush will delete all ExchangePayload for a peer.
func (s *SimpleExchangeContainer) Flush(id string) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.db, id)

}

// Values return all ExchangePayload for a peer.
func (s *SimpleExchangeContainer) Values(id string) []ExchangePayload {

	s.mutex.RLock()
	e, ok := s.db[id]
	s.mutex.RUnlock()

	if !ok {
		return []ExchangePayload{}
	}

	return e.Values()

}

// For concurrent access, a peer has it own container.
// It uses a Read/Write mutex.
type exchangePayloadContainer struct {
	mutex sync.RWMutex
	db    map[string]ExchangePayload
}

func (c *exchangePayloadContainer) Add(e ExchangePayload) {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(c.db) < ExchangePayloadLimit {
		c.db[e.ID] = e
	}

}

func (c *exchangePayloadContainer) Remove(e ExchangePayload) {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.db, e.ID)

}

func (c *exchangePayloadContainer) Values() []ExchangePayload {

	c.mutex.RLock()
	defer c.mutex.RUnlock()

	l := []ExchangePayload{}

	for _, e := range c.db {
		l = append(l, e)
	}

	return l

}
