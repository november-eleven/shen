package container

import (
	"encoding/json"
	"sync"
)

type ExchangePayload struct {
	ID  string          `json:"id"`
	Raw json.RawMessage `json:"signal"`
}

type ExchangeContainer interface {
	Add(id string, payload ExchangePayload)
	Remove(id string, payload ExchangePayload)
	Flush(id string)
	Values(id string) []ExchangePayload
}

type SimpleExchangeContainer struct {
	mutex sync.RWMutex
	db    map[string](*exchangePayloadContainer)
}

func NewExchangeContainer() ExchangeContainer {
	return &SimpleExchangeContainer{db: make(map[string](*exchangePayloadContainer))}

}

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

func (s *SimpleExchangeContainer) Remove(id string, e ExchangePayload) {

	s.mutex.Lock()
	l, ok := s.db[id]
	s.mutex.Unlock()

	if ok {
		l.Remove(e)
	}

}

func (s *SimpleExchangeContainer) Flush(id string) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.db, id)

}

func (s *SimpleExchangeContainer) Values(id string) []ExchangePayload {

	s.mutex.RLock()
	e, ok := s.db[id]
	s.mutex.RUnlock()

	if !ok {
		return []ExchangePayload{}
	}

	return e.Values()

}

type exchangePayloadContainer struct {
	mutex sync.RWMutex
	db    map[string]ExchangePayload
}

func (c *exchangePayloadContainer) Add(e ExchangePayload) {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.db[e.ID] = e

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
