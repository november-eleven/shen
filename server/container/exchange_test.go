package container

import (
	"testing"
)

func TestNewExchangeContainer(t *testing.T) {

	c := NewExchangeContainer()
	if c == nil {
		t.Fatal("Unexpected nil value for NewExchangeContainer")
	}

}

func TestSimpleExchangeContainerAdd(t *testing.T) {

	c := newSimpleExchangeContainer()

	var fixtures = []struct {
		id string
		p  ExchangePayload
	}{
		{"1770f3c9-baf8-4f0c-bfa8-79021abc164c", ExchangePayload{ID: "9a51b80b-ab1d-42f1-b471-05a723414ffa"}},
		{"32840865-8f9c-4627-a76e-d9b395446bbe", ExchangePayload{ID: "32840865-8f9c-4627-a76e-d9b395446bbe"}},
		{"9a51b80b-ab1d-42f1-b471-05a723414ffa", ExchangePayload{ID: "4774e326-dd90-4b07-8852-421f61ad4419"}},
	}

	for _, f := range fixtures {

		c.Add(f.id, f.p)

		m, ok := c.db[f.id]

		if !ok {
			t.Fatalf("unexpected empty row (first) for this fixture: %+v", f)
		}

		p, ok := m.db[f.p.ID]

		if !ok {
			t.Fatalf("unexpected empty row (last) for this fixture: %+v", f)
		}

		if p.ID != f.p.ID {
			t.Fatalf("stale write for this fixture: %+v", f)
		}

	}

}

func TestSimpleExchangeContainerRemove(t *testing.T) {

	c := newSimpleExchangeContainer()

	var fixtures = []struct {
		id string
		p  ExchangePayload
	}{
		{"1770f3c9-baf8-4f0c-bfa8-79021abc164c", ExchangePayload{ID: "9a51b80b-ab1d-42f1-b471-05a723414ffa"}},
		{"32840865-8f9c-4627-a76e-d9b395446bbe", ExchangePayload{ID: "32840865-8f9c-4627-a76e-d9b395446bbe"}},
		{"9a51b80b-ab1d-42f1-b471-05a723414ffa", ExchangePayload{ID: "4774e326-dd90-4b07-8852-421f61ad4419"}},
	}

	for _, f := range fixtures {
		c.Add(f.id, f.p)
	}

	for _, f := range fixtures {

		c.Remove(f.id, f.p)

		m, ok := c.db[f.id]

		if !ok {
			t.Fatalf("unexpected empty row (first) for this fixture: %+v", f)
		}

		p, ok := m.db[f.p.ID]

		if ok {
			t.Fatalf("unexpected result for this fixture: %+v", f)
		}

		if ok && p.ID != f.p.ID {
			t.Fatalf("stale write for this fixture: %+v", f)
		}

	}

}

func TestSimpleExchangeContainerFlush(t *testing.T) {

	c := newSimpleExchangeContainer()

	var fixtures = []struct {
		id string
		p  ExchangePayload
	}{
		{"1770f3c9-baf8-4f0c-bfa8-79021abc164c", ExchangePayload{ID: "9a51b80b-ab1d-42f1-b471-05a723414ffa"}},
		{"32840865-8f9c-4627-a76e-d9b395446bbe", ExchangePayload{ID: "32840865-8f9c-4627-a76e-d9b395446bbe"}},
		{"9a51b80b-ab1d-42f1-b471-05a723414ffa", ExchangePayload{ID: "4774e326-dd90-4b07-8852-421f61ad4419"}},
	}

	for _, f := range fixtures {
		c.Add(f.id, f.p)
	}

	for _, f := range fixtures {

		c.Flush(f.id)

		_, ok := c.db[f.id]

		if ok {
			t.Fatalf("unexpected result (first row) for this fixture: %+v", f)
		}

	}

}

func TestSimpleExchangeContainerValues(t *testing.T) {

	c := newSimpleExchangeContainer()

	var fixtures = []struct {
		id string
		p  ExchangePayload
	}{
		{"1770f3c9-baf8-4f0c-bfa8-79021abc164c", ExchangePayload{ID: "9a51b80b-ab1d-42f1-b471-05a723414ffa"}},
		{"32840865-8f9c-4627-a76e-d9b395446bbe", ExchangePayload{ID: "32840865-8f9c-4627-a76e-d9b395446bbe"}},
		{"9a51b80b-ab1d-42f1-b471-05a723414ffa", ExchangePayload{ID: "4774e326-dd90-4b07-8852-421f61ad4419"}},
	}

	for _, f := range fixtures {
		c.Add(f.id, f.p)
	}

	for _, f := range fixtures {

		l := c.Values(f.id)

		if len(l) != 1 {
			t.Fatalf("unexpected number of result for this fixture: %+v", f)
		}

		if l[0].ID != f.p.ID {
			t.Fatalf("stale read for this fixture: %+v", f)
		}

	}

}

func newSimpleExchangeContainer() *SimpleExchangeContainer {
	return &SimpleExchangeContainer{db: make(map[string](*exchangePayloadContainer))}
}
