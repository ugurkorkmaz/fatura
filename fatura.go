package fatura

import (
	"encoding/json"
	"errors"
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

var client *http.Client = &http.Client{}

type (
	Fatura interface {
		GetTestCredentials() (username, password string, err error)
		gateway(path Path) string
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

func (f *fatura) gateway(path Path) string {
	if f.debug {
		return string(TEST) + string(path)
	}
	return string(BASE) + string(path)
}
