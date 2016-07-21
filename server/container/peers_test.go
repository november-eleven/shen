package container

import (
	"testing"
)

func TestNewPeersContainer(t *testing.T) {

	c := NewPeersContainer()
	if c == nil {
		t.Fatal("Unexpected nil value for NewPeersContainer")
	}

}

func TestSimplePeersContainerAdd(t *testing.T) {

	c := newSimplePeersContainer()

	var fixtures = []string{
		"1770f3c9-baf8-4f0c-bfa8-79021abc164c",
		"32840865-8f9c-4627-a76e-d9b395446bbe",
		"9a51b80b-ab1d-42f1-b471-05a723414ffa",
	}

	for _, f := range fixtures {

		c.Add(f)

		m, ok := c.db[f]

		if !ok {
			t.Fatalf("unexpected empty result for this fixture: %+v", f)
		}

		if !m {
			t.Fatalf("stale write for this fixture: %+v", f)
		}

	}

}

func TestSimplePeersContainerRemove(t *testing.T) {

	c := newSimplePeersContainer()

	var fixtures = []string{
		"1770f3c9-baf8-4f0c-bfa8-79021abc164c",
		"32840865-8f9c-4627-a76e-d9b395446bbe",
		"9a51b80b-ab1d-42f1-b471-05a723414ffa",
	}

	for _, f := range fixtures {
		c.Add(f)
	}

	for _, f := range fixtures {

		c.Remove(f)

		_, ok := c.db[f]

		if ok {
			t.Fatalf("unexpected result for this fixture: %+v", f)
		}

	}

}

func TestSimplePeersContainerValues(t *testing.T) {

	c := newSimplePeersContainer()

	var fixtures = []string{
		"1770f3c9-baf8-4f0c-bfa8-79021abc164c",
		"32840865-8f9c-4627-a76e-d9b395446bbe",
		"9a51b80b-ab1d-42f1-b471-05a723414ffa",
	}

	for _, f := range fixtures {
		c.Add(f)
	}

	l := c.Values()

	if len(l) != len(fixtures) {
		t.Fatalf("unexpected result: %+v", l)
	}

loop:
	for _, f := range fixtures {
		for _, e := range l {
			if e == f {
				continue loop
			}
		}
		t.Fatalf("unexpected result: %+v", l)
	}

}

func newSimplePeersContainer() *SimplePeersContainer {
	return &SimplePeersContainer{db: make(map[string]bool)}
}
