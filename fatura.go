package fatura

import (
	"encoding/json"
	"errors"
	"fatura/entity"
	"fatura/entity/enum/document"
	"fmt"
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

type ApiError struct {
	Error    bool `json:"error"`
	Messages []struct {
		Type int64  `json:"type"`
		Text string `json:"text"`
	} `json:"messages"`
}

type (
	// Fatura interface.
	Fatura interface {
		// Login to the server.
		Login() error
		// Logout from the server.
		Logout() error
		// Gateway returns the gateway url.
		gateway(path Path) string
		// Get oid sms verification, step 1.
		StartSmsVerification(phone string) (string, error)
		// Get oid sms verification, step 2.
		EndSmsVerification(oid, code string, invocies []string) error
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
		SetCredentials(username, password string) Fatura
		// Update the user information on the server.
		UpdateUser(user *entity.User) (err error)
	}
	bearer struct {
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
	return &bearer{
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

// Gateway returns the gateway url.
func (b *bearer) gateway(path Path) string {
	if b.debug {
		return string(TEST) + string(path)
	}
	return string(BASE) + string(path)
}

// Get the token from the server.
func (b *bearer) GetToken() string {
	return b.token
}

// Set the debug mode.
func (b *bearer) SetDebug(debug bool) Fatura {
	b.debug = debug
	return b
}

// Get the debug mode.
func (b *bearer) GetDebug() bool {
	return b.debug
}

// Set the username and password.
func (b *bearer) SetCredentials(username, password string) Fatura {
	b.username = username
	b.password = password
	return b
}

// Get the username and password.
func (b *bearer) GetCridetials() (username, password string) {
	return b.username, b.password
}

// Get the test credentials from the server.
func (b *bearer) GetTestCredentials() (username, password string, err error) {
	res, err := client.PostForm(b.gateway(ESIGN), url.Values{
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
func (b *bearer) Login() error {
	if b.username == "" || b.password == "" {
		return errors.New("username or password is empty")
	}
	assoscmd := []string{"anologin"}
	if b.debug {
		assoscmd = []string{"login"}
	}
	res, err := client.PostForm(b.gateway(LOGIN), url.Values{
		"assoscmd": assoscmd,
		"userid":   []string{b.username},
		"sifre":    []string{b.password},
		"sifre2":   []string{b.password},
		"parola":   []string{b.password},
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
		fmt.Println(b.token)
		fmt.Println(string(jsonData))
		return errors.New("token is nil")
	}
	b.token = data["token"].(string)
	return nil
}

// Logout from the server.
func (b *bearer) Logout() error {
	if b.token == "" {
		return errors.New("token is empty")
	}
	res, err := client.PostForm(b.gateway(LOGIN), url.Values{
		"assoscmd": []string{"logout"},
		"token":    []string{b.token},
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
	b.username = ""
	b.password = ""
	b.token = ""
	return nil
}

// Get the user information from the server.
func (b *bearer) GetUser() (user *entity.User, err error) {
	res, err := client.PostForm(b.gateway(DISPATCH), url.Values{
		"callid":   []string{b.uuid.String()},
		"token":    []string{b.token},
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

// Update the user information on the server.
func (b *bearer) UpdateUser(user *entity.User) error {
	res, err := client.PostForm(b.gateway(DISPATCH), url.Values{
		"callid":   []string{b.uuid.String()},
		"token":    []string{b.token},
		"cmd":      []string{"EARSIV_PORTAL_KULLANICI_BILGILERI_GETIR"},
		"pageName": []string{"RG_KULLANICI"},
		"jp":       []string{""},
	})
	if err != nil {
		return errors.New("Error while sending request: " + err.Error())
	}
	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New("Error while reading response: " + err.Error())
	}
	var result struct {
		Data entity.User `json:"data"`
	}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return errors.New("Error while parsing response: " + err.Error())
	}
	return nil
}

// Starts the sms verification process.
func (b *bearer) StartSmsVerification(phone string) (string, error) {
	var err error
	var jp struct {
		CEPTEL  string `json:"CEPTEL"`
		KCEPTEL string `json:"KCEPTEL"`
		TIP     string `json:"TIP"`
	}
	jp.CEPTEL = phone
	jp.KCEPTEL = ""
	jp.TIP = "1"

	jsonData, err := json.Marshal(jp)
	if err != nil {
		return "", errors.New("Error while parsing data: " + err.Error())
	}
	res, err := client.PostForm(b.gateway(DISPATCH), url.Values{
		"callid":   []string{uuid.NewString()},
		"token":    []string{b.token},
		"cmd":      []string{"EARSIV_PORTAL_SMSSIFRE_GONDER"},
		"pageName": []string{"RG_SMSONAY"},
		"jp":       []string{string(jsonData)},
	})
	if err != nil {
		return "", errors.New("Error while sending request: " + err.Error())
	}
	jsonData, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("Error while reading response: " + err.Error())
	}

	var result struct {
		Data struct {
			Oid string `json:"oid"`
		} `json:"data"`
	}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return "", errors.New("Error while parsing response: " + err.Error())
	}

	if result.Data.Oid == "" {
		return "", errors.New("oid is empty")
	}
	return result.Data.Oid, nil
}

func (b *bearer) EndSmsVerification(oid, code string, uuids []string) error {
	params := map[string]interface{}{
		"DATA":  uuids,
		"SIFRE": code,
		"OID":   oid,
		"OPR":   1,
	}

	_json, err := json.Marshal(params)
	if err != nil {
		return errors.New("Error while parsing data: " + err.Error())
	}
	res, err := client.PostForm(b.gateway(DISPATCH), url.Values{
		"callid":   []string{uuid.NewString()},
		"token":    []string{b.token},
		"cmd":      []string{"0lhozfib5410mp"},
		"pageName": []string{"RG_SMSONAY"},
		"jp":       []string{string(_json)},
	})
	if err != nil {
		return errors.New("Error while sending request: " + err.Error())
	}
	_data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New("Error while reading response: " + err.Error())
	}

	var _err ApiError = ApiError{}
	err = json.Unmarshal(_data, &_err)
	if err != nil {
		return errors.New("Error while parsing response: " + err.Error())
	}
	for _, e := range _err.Messages {
		return errors.New(e.Text)
	}

	var result struct {
		Data struct {
			Sonuc string `json:"sonuc"`
		} `json:"data"`
	}
	err = json.Unmarshal(_data, &result)
	if err != nil {
		return errors.New("Error while parsing response: " + err.Error())
	}

	if result.Data.Sonuc != "1" {
		return errors.New("error while verifying sms code")
	}
	for i := range uuids {
		b.rowCount = i + 1
	}

	return nil
}
