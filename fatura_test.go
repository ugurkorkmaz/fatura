package fatura_test

import (
	"fatura"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fatura_new = fatura.New()

func TestSetDebug(t *testing.T) {
	fatura_new.SetDebug(true)
	if !fatura_new.GetDebug() {
		t.Error("SetDebug() failed")
	}
}

func TestGetDebug(t *testing.T) {
	fatura_new.SetDebug(true)
	if !fatura_new.GetDebug() {
		t.Error("GetDebug() failed")
	}
}

func TestSetCridetials(t *testing.T) {
	f := fatura_new.SetCridetials("username", "password")

	username, password := f.GetCridetials()
	assert.Equal(t, "username", username)
	assert.Equal(t, "password", password)
}

func TestGetCridetials(t *testing.T) {
	f := fatura_new.SetCridetials("username", "password")
	username, password := f.GetCridetials()
	assert.Equal(t, "username", username)
	assert.Equal(t, "password", password)
}

func TestGetTestCredentials(t *testing.T) {
	_, password, err := fatura_new.GetTestCredentials()
	if err != nil {
		t.Error("GetTestCredentials() failed")
	}
	assert.Equal(t, "1", password)
}

func TestLogin(t *testing.T) {
	username, password, err := fatura_new.GetTestCredentials()
	if err != nil {
		t.Error("GetTestCredentials() failed")
	}

	f := fatura_new.SetCridetials(username, password)
	f.SetDebug(true)
	err = f.Login()
	if err != nil {
		t.Error("Login() failed")
	}
}

func TestLogout(t *testing.T) {
	username, password, err := fatura_new.GetTestCredentials()
	if err != nil {
		t.Error("GetTestCredentials() failed")
	}

	f := fatura_new.SetCridetials(username, password)
	f.SetDebug(true)
	err = f.Login()
	if err != nil {
		t.Error("Login() failed")
	}

	err = f.Logout()
	if err != nil {
		t.Error("Logout() failed")
	}
}
