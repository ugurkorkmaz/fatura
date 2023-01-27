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
	"strings"

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

/*
ApiError is the error returned by the server.
*/
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
		/*
			Login to the server.
		*/
		Login() error
		/*
			Logout from the server.
		*/
		Logout() error
		/*
			Gateway returns the gateway url.
		*/
		gateway(path Path) string
		/*
			Get oid sms verification, step 1.
		*/
		StartSmsVerification(phone string) (string, error)
		/*
			Get oid sms verification, step 2.
		*/
		EndSmsVerification(oid, code string, invocies []string) error
		/*
		 Creates a draft.
		 The model parameter can be one of the following:

		 * entity.Invoice

		 * entity.ProducerReceipt

		 * entity.SelfEmployedReceipt
		*/
		CreateDraft(entity any) error
		/*
			Delete a draft.
		*/
		DeleteDraft(document []string, reasons string) error
		/*
			Extends the getter interface.
		*/
		getter
		/*
			Extends the setter interface.
		*/
		setter
		/*
			Extends the lister interface.
		*/
		lister
	}
	// Getter interface.
	getter interface {
		/*
			Get the token from the server.
		*/
		GetToken() string
		/*
			Get the token from the server.
		*/
		GetTestCredentials() (username, password string, err error)
		/*
			Get the username and password.
		*/
		GetCridetials() (username, password string)
		/*
			Get the debug mode.
		*/
		GetDebug() bool
		/*
			Get the user information from the server.
		*/
		GetUser() (user *entity.User, err error)
		/*
			Get the document download url.
		*/
		GetDownloadURL(id uuid.UUID, signed bool) (string, error)
		/*
		 Get the document html content.
		*/
		GetHtml(id uuid.UUID, signed bool) ([]byte, error)
	}
	// Setter interface.
	setter interface {
		/*
			Set the debug mode.
		*/
		SetDebug(bool) Fatura
		/*
			Set the username and password.
		*/
		SetCredentials(username, password string) Fatura
		/*
			Update the user information on the server.
		*/
		UpdateUser(user *entity.User) (err error)
	}
	lister interface {
		/*
			Cancel a request.
		*/
		CancellationRequest(uuid.UUID, string) string
		/*
			Objection a request.
		*/
		ObjectionRequest() string
		/*
			Objection a request.
		*/
		GetRequests(start, end string) []string
		/*
			Get all documents.
		*/
		GetAll(start, end string)
		/*
			Filters documents.
		*/
		GetAllIssuedToMe(start, end, hourlySearch string)
		/*
			Filters documents.
		*/
		FilterDocuments(document.Type)
		/*
			Select a column.
		*/
		SelectColumn(column, key string) string
		/*
			Map a column.
		*/
		MapColumn(data []string) entity.Array
		/*
			Set the filters.
		*/
		SetFilters(filters []string) lister
		/*
			Set the limit.
		*/
		SetLimit(limit, offset int) lister
		/*
			Ascending sort.
		*/
		SortAsc() lister
		/*
			Descending sort.
		*/
		SortDesc() lister
		/*
			Set the row count.
		*/
		SetRowCount(int) lister
		/*
			Get the row count.
		*/
		RowCount() int
		/*
			Only signed documents.
		*/
		OnlySigned() lister
		/*
			Only unsigned documents.
		*/
		OnlyUnsigned() lister
		/*
			Only deleted documents.
		*/
		OnlyDeleted() lister
		/*
			Only active documents.
		*/
		OnlyCurrent() lister
		/*
			Only invoice documents.
		*/
		OnlyInvoice() lister
		/*
			Only producer receipt documents.
		*/
		OnlyProducerReceipt() lister
		/*
			Only self employed receipt documents.
		*/
		OnlySelfEmployedReceipt() lister
		/*
			Search by recipient name.
		*/
		FindRecipientName(string) lister
		/*
			Search by recipient id.
		*/
		FindRecipientId(string) lister
		/*
			Search by recipient tax number.
		*/
		FindEttn(string) lister
		/*
			Search by document id.
		*/
		FindDocumentId(string) lister
	}
)

// Fatura is the main interface.
type bearer struct {
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

// Ends the sms verification process.
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

/*
Creates a draft.
The model parameter can be one of the following:

* entity.Invoice

* entity.ProducerReceipt

* entity.SelfEmployedReceipt
*/
func (b *bearer) CreateDraft(model any) error {
	var form url.Values

	switch model := model.(type) {
	case *entity.Invoice:
		form.Add("jp", model.Json())
		form.Add("cmd", "EARSIV_PORTAL_FATURA_OLUSTUR")
		form.Add("pageName", "RG_BASITFATURA")
	case *entity.ProducerReceipt:
		form.Add("jp", model.Json())
		form.Add("cmd", "EARSIV_PORTAL_MUSTAHSIL_OLUSTUR")
		form.Add("pageName", "RG_MUSTAHSIL")
	case *entity.SelfEmployedReceipt:
		form.Add("jp", model.Json())
		form.Add("cmd", "EARSIV_PORTAL_SERBEST_MESLEK_MAKBUZU_OLUSTUR")
		form.Add("pageName", "RG_SERBEST")
	default:
		return errors.New("invalid model")
	}

	res, err := client.PostForm(b.gateway(DISPATCH), form)
	if err != nil {
		return errors.New("Error while sending request: " + err.Error())
	}
	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New("Error while reading response: " + err.Error())
	}

	var result ApiError = ApiError{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return errors.New("Error while parsing response: " + err.Error())
	}
	if len(result.Messages) > 0 {
		for _, e := range result.Messages {
			return errors.New(e.Text)
		}
	}
	if !strings.Contains(string(jsonData), "başarıyla") {
		return errors.New("error while creating draft")
	}
	return nil
}

func (b *bearer) DeleteDraft(uuids []string, reasons string) error {
	var jp []map[string]interface{}
	for _, uuid := range uuids {
		jp = append(jp, map[string]interface{}{
			"belgeTuru": b.document.String(),
			"ettn":      uuid,
		})
	}
	_json, err := json.Marshal(jp)
	if err != nil {
		return errors.New("Error while parsing data: " + err.Error())
	}
	res, err := client.PostForm(b.gateway(DISPATCH), url.Values{
		"token":    []string{b.token},
		"cmd":      []string{"EARSIV_PORTAL_FATURA_SIL"},
		"pageName": []string{"RG_TASLAKLAR"},
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

	return nil
}

// Set the document type.
func (b *bearer) SetDocumentType(document document.Type) Fatura {
	b.document = document
	return b
}

// Get the list of invoices.
func (b *bearer) GetDocument(id uuid.UUID) (*entity.Array, error) {
	params := url.Values{}
	switch b.document {
	case document.Invoice:
		params.Add("cmd", "EARSIV_PORTAL_FATURA_GETIR")
		params.Add("pageName", "RG_TASLAKLAR")
	case document.ProducerReceipt:
		params.Add("cmd", "EARSIV_PORTAL_MUSTAHSIL_GETIR")
		params.Add("pageName", "RG_MUSTAHSIL")
	case document.SelfEmployedReceipt:
		params.Add("cmd", "EARSIV_PORTAL_SERBEST_MESLEK_GETIR")
		params.Add("pageName", "RG_SERBEST")
	}
	params.Add("ettn", id.String())

	res, err := client.PostForm(b.gateway(DISPATCH), params)
	if err != nil {
		return nil, errors.New("Error while sending request: " + err.Error())
	}
	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Error while reading response: " + err.Error())
	}
	var result = &entity.Array{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, errors.New("Error while parsing response: " + err.Error())
	}
	return result, nil
}

// Get the download url of the document.
func (b *bearer) GetDownloadURL(id uuid.UUID, signed bool) (string, error) {
	res, err := client.PostForm(b.gateway(DISPATCH), url.Values{
		"token":    []string{b.token},
		"ettn":     []string{id.String()},
		"cmd":      []string{"EARSIV_PORTAL_FATURA_GOSTER"},
		"pageName": []string{"RG_TASLAKLAR"},
		"onayDurumu": func() []string {
			if signed {
				return []string{"Onaylandı"}
			}
			return []string{"Onaylanmadı"}
		}(),
	})

	if err != nil {
		return "", errors.New("Error while sending request: " + err.Error())
	}
	jsonData, err := ioutil.ReadAll(res.Body)
	return string(jsonData), nil

}

// Get the html of the document.
func (b *bearer) GetHtml(id uuid.UUID, signed bool) ([]byte, error) {
	return nil, errors.New("not implemented")
}

// TODO
func (b *bearer) CancellationRequest(id uuid.UUID, explanation string) string {
	res, err := client.PostForm(b.gateway(DISPATCH), url.Values{
		"token":         []string{b.token},
		"ettn":          []string{id.String()},
		"cmd":           []string{"EARSIV_PORTAL_IPTAL_TALEBI_OLUSTUR"},
		"pageName":      []string{"RG_TASLAKLAR"},
		"belgeTuru":     []string{b.document.String()},
		"talepAciklama": []string{explanation},
	})
	if err != nil {
		return ""
	}
	jsonData, err := ioutil.ReadAll(res.Body)
	return string(jsonData)
}

// TODO
func (b *bearer) ObjectionRequest() string {
	panic("not implemented")
}

// TODO
func (b *bearer) GetRequests(start, end string) []string {
	panic("not implemented")
}

// TODO
func (b *bearer) GetAll(start, end string) {
	panic("not implemented")
}

// TODO
func (b *bearer) GetAllIssuedToMe(start, end, hourlySearch string) {
	panic("not implemented")
}

// TODO
func (b *bearer) FilterDocuments(document.Type) {
	panic("not implemented")
}

// TODO
func (b *bearer) SelectColumn(column, key string) string {
	panic("not implemented")
}

// TODO
func (b *bearer) MapColumn(data []string) entity.Array {
	panic("not implemented")
}

// TODO
func (b *bearer) SetFilters(filters []string) lister {
	panic("not implemented")
}

// TODO
func (b *bearer) SetLimit(limit, offset int) lister {
	panic("not implemented")
}

// Sort the list in ascending order.
func (b *bearer) SortAsc() lister {
	b.sortByDesc = false
	return b
}

// Sort the list in descending order.
func (b *bearer) SortDesc() lister {
	b.sortByDesc = true
	return b
}

// Set the row count.
func (b *bearer) SetRowCount(count int) lister {
	b.rowCount = count
	return b
}

// Get the row count.
func (b *bearer) RowCount() int {
	return b.rowCount
}

// TODO
func (b *bearer) OnlySigned() lister {
	panic("not implemented")
}

// TODO
func (b *bearer) OnlyUnsigned() lister {
	panic("not implemented")
}

// TODO
func (b *bearer) OnlyDeleted() lister {
	panic("not implemented")
}

// TODO
func (b *bearer) OnlyCurrent() lister {
	panic("not implemented")
}

// TODO
func (b *bearer) OnlyInvoice() lister {
	panic("not implemented")
}

// TODO
func (b *bearer) OnlyProducerReceipt() lister {
	panic("not implemented")
}

// TODO
func (b *bearer) OnlySelfEmployedReceipt() lister {
	panic("not implemented")
}

// TODO
func (b *bearer) FindRecipientName(string) lister {
	panic("not implemented")
}

// TODO
func (b *bearer) FindRecipientId(string) lister {
	panic("not implemented")
}

// TODO
func (b *bearer) FindEttn(string) lister {
	panic("not implemented")
}

// TODO
func (b *bearer) FindDocumentId(string) lister {
	panic("not implemented")
}
