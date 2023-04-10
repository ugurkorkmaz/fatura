package fatura_test

import (
	"fatura"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fatura_new = fatura.New()

var username, password string

func init() {
	var err error
	username, password, err = fatura_new.GetTestCredentials()
	if err != nil {
		panic(err)
	}
	fatura_new = fatura_new.Debug()
}

func TestGetTestCredentials(t *testing.T) {
	_, password, err := fatura_new.GetTestCredentials()
	if err != nil {
		t.Error("GetTestCredentials() failed")
	}
	assert.Equal(t, "1", password)
}

func TestLogin(t *testing.T) {
	u, err := fatura_new.Login(username, password)
	if err != nil {
		t.Error("Login() failed")
	}
	if u == nil {
		t.Error("Login() failed")
	}

	if u.Token() == "" {
		t.Error("Login() failed")
	}
}

func TestLogout(t *testing.T) {
	_, err := fatura_new.Login(username, password)
	if err != nil {
		t.Error("Login() failed")
	}
	err = fatura_new.Logout()
	if err != nil {
		t.Error("Logout() failed")
	}
}

func TestGetUser(t *testing.T) {
	_, err := fatura_new.Login(username, password)
	if err != nil {
		t.Error("Login() failed")
	}

	u, err := fatura_new.User()
	if err != nil {
		t.Error("GetUser() failed")
	}
	if u == nil {
		t.Error("GetUser() failed")
	}
	err = fatura_new.Logout()
	if err != nil {
		t.Error("Logout() failed")
	}
}
