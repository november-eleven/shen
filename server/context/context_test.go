package context

import (
	"testing"

	"github.com/november-eleven/shen/server/container"
)

func TestNewRepository(t *testing.T) {

	c := NewRepository(container.NewExchangeContainer(), container.NewPeersContainer())
	if c == nil {
		t.Fatal("Unexpected nil value for NewRepository")
	}

}

func TestLogin(t *testing.T) {

	c := NewRepository(container.NewExchangeContainer(), container.NewPeersContainer())

	if c.Login() == "" {
		t.Fatal("Unexpected empty result for Login")
	}

}

func TestLogout(t *testing.T) {

	c := NewRepository(container.NewExchangeContainer(), container.NewPeersContainer())
	l := c.Login()
	c.Logout(l)

}
