package fatura_test

import (
	"fatura"
	"fatura/entity"
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

func TestSetCredentials(t *testing.T) {
	f := fatura_new.SetCredentials("username", "password")

	username, password := f.GetCridetials()
	assert.Equal(t, "username", username)
	assert.Equal(t, "password", password)
}

func TestGetCridetials(t *testing.T) {
	f := fatura_new.SetCredentials("username", "password")
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

	f := fatura_new.SetCredentials(username, password)
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

	f := fatura_new.SetCredentials(username, password)
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

func TestGetUser(t *testing.T) {
	username, password, err := fatura_new.GetTestCredentials()
	if err != nil {
		t.Error("GetTestCredentials() failed")
	}

	f := fatura_new.SetCredentials(username, password)
	f.SetDebug(true)
	err = f.Login()
	if err != nil {
		t.Error("Login() failed")
	}

	_, err = f.GetUser()
	if err != nil {
		t.Error("GetUser() failed")
	}
	f.Logout()
}

func TestUpdateUser(t *testing.T) {
	username, password, err := fatura_new.GetTestCredentials()
	if err != nil {
		t.Error("GetTestCredentials() failed")
	}

	f := fatura_new.SetCredentials(username, password)
	f.SetDebug(true)
	err = f.Login()
	if err != nil {
		t.Error("Login() failed")
	}

	user := &entity.User{}
	user.Kasaba = "Kasaba"

	err = f.UpdateUser(user)
	if err != nil {
		t.Error("UpdateUser() failed")
	}

	user, err = f.GetUser()
	if err != nil {
		t.Error("GetUser() failed")
	}

	assert.Equal(t, "Kasaba", user.Kasaba)
	err = f.Logout()
	if err != nil {
		t.Error("Logout() failed")
	}
}
