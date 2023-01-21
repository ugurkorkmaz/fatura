package fatura

import (
	"encoding/json"
	"errors"
	"fatura/entity"
	"fatura/entity/enum/document"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type (
	Api  string
	Path string
)

const (
	BASE Api = "https://earsivportal.efatura.gov.tr"
	TEST Api = "https://earsivportaltest.efatura.gov.tr"
)

const (
	DISPATCH Path = "/earsiv-services/dispatch"
	DOWNLOAD Path = "/earsiv-services/download"
	LOGIN    Path = "/earsiv-services/assos-login"
	REFERRER Path = "/intragiris.html"
	ESIGN    Path = "/earsiv-services/esign"
)

// HTTP client.
var client *http.Client = &http.Client{}

// Fatura instance.
type (
	// Fatura interface.
	Fatura interface {
		// Login to the server.
		Login() error
		// Logout from the server.
		Logout() error
		// Gateway returns the gateway url.
		gateway(path Path) string
		// Extends the getter interface.
		getter
		// Extends the setter interface.
		setter
	}
	// Getter interface.
	getter interface {
		// Get the token from the server.
		GetToken() string
		// Get the token from the server.
		GetTestCredentials() (username, password string, err error)
		// Get the username and password.
		GetCridetials() (username, password string)
		// Get the debug mode.
		GetDebug() bool
		// Get the user information from the server.
		GetUser() (user *entity.User, err error)
	}
	// Setter interface.
	setter interface {
		// Set the debug mode.
		SetDebug(bool) Fatura
		// Set the username and password.
		SetCridetials(username, password string) Fatura
	}
	fatura struct {
		uuid       uuid.UUID
		sortByDesc bool
		rowCount   int
		column     []string
		filters    []string
		limit      []int
		document   document.Type
		debug      bool
		username   string
		password   string
		token      string
	}
)

// Returns a new Fatura instance.
func New() Fatura {
	return &fatura{
		uuid:       uuid.New(),
		sortByDesc: false,
		rowCount:   100,
		column:     []string{"id", "uuid", "type", "status", "date", "amount"},
		filters:    []string{},
		limit:      []int{0, 100},
		document:   document.Invoice,
		debug:      false,
	}
}

// Get the token from the server.
func (f *fatura) GetToken() string {
	return f.token
}

// Set the debug mode.
func (f *fatura) SetDebug(debug bool) Fatura {
	f.debug = debug
	return f
}

// Get the debug mode.
func (f *fatura) GetDebug() bool {
	return f.debug
}

// Set the username and password.
func (f *fatura) SetCridetials(username, password string) Fatura {
	f.username = username
	f.password = password
	return f
}

// Get the username and password.
func (f *fatura) GetCridetials() (username, password string) {
	return f.username, f.password
}

// Get the test credentials from the server.
func (f *fatura) GetTestCredentials() (username, password string, err error) {
	res, err := client.PostForm(f.gateway(ESIGN), url.Values{
		"assoscmd": []string{"kullaniciOner"},
		"rtype":    []string{"json"},
	})

	if err != nil {
		return "", "", errors.New("Error while sending request: " + err.Error())
	}

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", "", errors.New("Error while reading response: " + err.Error())
	}

	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return "", "", errors.New("Error while parsing response: " + err.Error())
	}
	if data["userid"].(string) == "" {
		return "", "", errors.New("Error while parsing response: " + err.Error())
	}
	return data["userid"].(string), "1", nil
}

// Login to the server.
func (f *fatura) Login() error {
	if f.username == "" || f.password == "" {
		return errors.New("username or password is empty")
	}
	assoscmd := []string{"anologin"}
	if f.debug {
		assoscmd = []string{"login"}
	}
	res, err := client.PostForm(f.gateway(LOGIN), url.Values{
		"assoscmd": assoscmd,
		"userid":   []string{f.username},
		"sifre":    []string{f.password},
		"sifre2":   []string{f.password},
		"parola":   []string{f.password},
	})
	if err != nil {
		return errors.New("Error while sending request: " + err.Error())
	}

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New("Error while reading response: " + err.Error())
	}

	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return errors.New("Error while parsing response: " + err.Error())
	}

	if data["token"] == nil {
		return errors.New("all credentials are wrong")
	}
	f.token = data["token"].(string)
	return nil
}

// Logout from the server.
func (f *fatura) Logout() error {
	if f.token == "" {
		return errors.New("token is empty")
	}
	res, err := client.PostForm(f.gateway(LOGIN), url.Values{
		"assoscmd": []string{"logout"},
		"token":    []string{f.token},
	})
	if err != nil {
		return errors.New("Error while sending request: " + err.Error())
	}

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New("Error while reading response: " + err.Error())
	}

	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return errors.New("Error while parsing response: " + err.Error())
	}
	f.username = ""
	f.password = ""
	f.token = ""
	return nil

}

// Get the user information from the server.
func (f *fatura) GetUser() (user *entity.User, err error) {
	res, err := client.PostForm(f.gateway(DISPATCH), url.Values{
		"callid":   []string{f.uuid.String()},
		"token":    []string{f.token},
		"cmd":      []string{"EARSIV_PORTAL_KULLANICI_BILGILERI_GETIR"},
		"pageName": []string{"RG_KULLANICI"},
		"jp":       []string{""},
	})
	if err != nil {
		return nil, errors.New("Error while sending request: " + err.Error())
	}
	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Error while reading response: " + err.Error())
	}
	var result struct {
		Data entity.User `json:"data"`
	}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, errors.New("Error while parsing response: " + err.Error())
	}
	return &result.Data, nil
}

func (f *fatura) gateway(path Path) string {
	if f.debug {
		return string(TEST) + string(path)
	}
	return string(BASE) + string(path)
}
