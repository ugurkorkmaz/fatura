package fatura

import (
	"bytes"
	"encoding/json"
	"errors"
	"fatura/entity"
	"fatura/entity/enum/document"
	"fatura/types"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
)

type Ifatura interface {
	// Login to the server.
	Login() error

	// Logout from the server.
	Logout() error

	// Gateway returns regular url.
	gateway(path Path) string

	// Get oid sms verification, step 1.
	StartSmsVerification(phone string) (string, error)

	// Get oid sms verification, step 2.
	EndSmsVerification(oid, code string, invocies []string) error

	/*
	 Creates a draft.
	 The model parameter can be one of the following:

	 * entity.Invoice

	 * entity.ProducerReceipt

	 * entity.SelfEmployedReceipt
	*/
	CreateDraft(entity any) error

	// Delete a draft.
	DeleteDraft(document []string, reasons string) error

	// Cancel a request.
	CancellationRequest(uuid.UUID, string) string

	// Objection a request.
	ObjectionRequest() string

	// Objection a request.
	GetRequests(start, end string) []string

	// Extends the getter interface.
	getter

	// Extends the setter interface.
	setter

	// Extends the lister interface.
	List() lister
}

type Fatura struct {
	uuid       uuid.UUID
	sortByDesc bool
	rowCount   int
	column     types.Array
	filters    types.Array
	limit      []int
	document   document.Type
	debug      bool
	username   string
	password   string
	token      string
}

func toPath(path Path, debug bool) string {
	baseMode := string(BASE)
	basePath := string(path)
	if debug {
		baseMode = string(TEST)
	}
	return baseMode + basePath
}

func gateway(path Path, debug bool, data types.Array) (types.Array, error) {
	var (
		address  = toPath(path, false)
		body     = url.Values(data)
		response = make(types.Array)
	)

	query, err := client.PostForm(address, body)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(query.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// This function returns a new instance of the Fatura.
func New() *Fatura {
	return &Fatura{
		uuid:       uuid.New(),
		sortByDesc: false,
		rowCount:   100,
		column:     types.Array{},
		filters:    types.Array{},
		limit:      []int{0, 100},
		document:   document.Invoice,
		debug:      false,
	}
}

// This function sets the debug flag to true.
func (f *Fatura) Debug() *Fatura {
	f.debug = true
	return f
}
func (f *Fatura) GetTestCredentials() (username, password string, err error) {

	var body = types.Array{}

	body.Add("assoscmd", "kullaniciOner")
	body.Add("rtype", "json")

	response, err := gateway(LOGIN, f.debug, body)
	if err != nil {
		return "", "", err
	}
	if response.Has("error") {
		return "", "", errors.New("login failed")
	}
	return response.Get("userid"), "1", nil
}

// Set the username and password.
func (f *Fatura) Login(username, password string) (fatura *Fatura, err error) {
	f.username = username
	f.password = password

	if f.username == "" || f.password == "" {
		return f, errors.New("username and password cannot be empty")
	}

	// Array type is a map[string]interface{}.
	var body = types.Array{}

	body.Add("assoscmd", "anologin")
	body.Add("userid", f.username)
	body.Add("sifre", f.password)
	body.Add("sifre2", f.password)
	body.Add("parola", f.password)

	// Check if debug mode is enabled.
	if f.debug {
		body.Set("assoscmd", "login")
	}

	// Send the request to the gateway.
	response, err := gateway(LOGIN, f.debug, body)
	if err != nil {
		return f, err
	}

	if response.Has("error") {
		return f, errors.New("login failed")
	}

	f.token = response.Get("token")
	return f, nil
}

func (f *Fatura) Logout() error {

	var body = types.Array{}

	if f.token == "" {
		return errors.New("token cannot be empty")
	}

	body.Add("assoscmd", "logout")
	body.Add("token", f.token)

	response, err := gateway(LOGIN, f.debug, body)
	if err != nil {
		return err
	}

	if response.Has("error") {
		return errors.New("logout failed")
	}

	f.token = ""
	f.username = ""
	f.password = ""
	return nil
}

// Get the token.
func (f *Fatura) Token() (token string) {
	return f.token
}

// Get the User.
func (f *Fatura) User() (*entity.User, error) {

	var body = types.Array{}

	body.Add("callid", f.uuid.String())
	body.Add("token", f.token)
	body.Add("cmd", "EARSIV_PORTAL_KULLANICI_BILGILERI_GETIR")
	body.Add("pageName", "RG_KULLANICI")
	body.Add("jp", "")

	response, err := gateway(DISPATCH, f.debug, body)
	if err != nil {
		return nil, err
	}

	if response.Has("error") {
		return nil, errors.New("user not found")
	}

	var user = entity.User{}

	err = json.Unmarshal([]byte(response.Json()), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

// Start the sms verification.
func (f *Fatura) StartSmsVerification(phone string) (string, error) {
	var (
		body = types.Array{}
		jp   = types.Array{}
	)
	jp.Add("CEPTEL", phone)
	jp.Add("KCEPTEL", "")
	jp.Add("TIP", "1")

	body.Add("callid", f.uuid.String())
	body.Add("token", f.token)
	body.Add("cmd", "EARSIV_PORTAL_SMSSIFRE_GONDER")
	body.Add("pageName", "RG_SMSONAY")
	body.Add("jp", jp.Json())

	response, err := gateway(DISPATCH, f.debug, body)
	if err != nil {
		return "", err
	}

	if response.Has("error") {
		return "", errors.New("sms verification failed")
	}

	if !response.Has("oid") {
		return "", errors.New("sms verification failed")
	}

	return response.Get("oid"), nil
}

// EndSmsVerification ends the sms verification.
func (f *Fatura) EndSmsVerification(oid, code string, uuids []string) error {
	var body = types.Array{}
	var jp = types.Array{}

	jp.Add("DATA", strings.Join(uuids, ","))
	jp.Add("SIFRE", code)
	jp.Add("OID", oid)
	jp.Add("TIP", "1")

	body.Add("callid", f.uuid.String())
	body.Add("token", f.token)
	body.Add("cmd", "0lhozfib5410mp")
	body.Add("pageName", "RG_SMSONAY")
	body.Add("jp", jp.Json())

	response, err := gateway(DISPATCH, f.debug, body)
	if err != nil {
		return err
	}

	if response.Has("error") {
		return errors.New("sms verification failed")
	}

	if !response.Has("oid") {
		return errors.New("sms verification failed")
	}

	if response.Get("sonuc") != "1" {
		return errors.New("sms verification failed")
	}

	return nil
}

func (f *Fatura) CreateDraft(model any) error {
	var body = types.Array{}

	switch model := model.(type) {
	case *entity.Invoice:
		body.Add("jp", model.Json())
		body.Add("cmd", "EARSIV_PORTAL_FATURA_OLUSTUR")
		body.Add("pageName", "RG_BASITFATURA")
	case *entity.ProducerReceipt:
		body.Add("jp", model.Json())
		body.Add("cmd", "EARSIV_PORTAL_MUSTAHSIL_OLUSTUR")
		body.Add("pageName", "RG_MUSTAHSIL")
	case *entity.SelfEmployedReceipt:
		body.Add("jp", model.Json())
		body.Add("cmd", "EARSIV_PORTAL_SERBEST_MESLEK_MAKBUZU_OLUSTUR")
		body.Add("pageName", "RG_SERBEST")
	default:
		return errors.New("invalid entity type")
	}

	response, err := gateway(DISPATCH, f.debug, body)
	if err != nil {
		return err
	}

	if response.Has("error") {
		return errors.New("draft creation failed")
	}

	if strings.Contains(response.Json(), "başarıyla") {
		return errors.New("draft creation failed")
	}
	return nil
}

func (f *Fatura) DeleteDraft(uuids []string, reasons string) error {
	var (
		body = types.Array{}
		jp   = types.Array{}
	)
	for _, uuid := range uuids {
		jp.Add("belgeTuru", "FATURA")
		jp.Add("ettn", uuid)
	}

	body.Add("callid", f.uuid.String())
	body.Add("token", f.token)
	body.Add("cmd", "EARSIV_PORTAL_FATURA_SIL")
	body.Add("pageName", "RG_BASITFATURA")
	body.Add("jp", jp.Json())

	response, err := gateway(DISPATCH, f.debug, body)
	if err != nil {
		return err
	}

	if response.Has("error") {
		return errors.New("draft deletion failed")
	}

	if strings.Contains(response.Json(), "başarıyla") {
		return errors.New("draft deletion failed")
	}
	return nil
}

// SetDocumentType sets the document type.
func (f *Fatura) SetDocumentType(document document.Type) Fatura {
	f.document = document
	return *f
}

func (f *Fatura) GetDocument(id uuid.UUID) (types.Array, error) {
	var body = types.Array{}

	switch f.document {
	case document.Invoice:
		body.Add("cmd", "EARSIV_PORTAL_FATURA_GETIR")
		body.Add("pageName", "RG_TASLAKLAR")
	case document.ProducerReceipt:
		body.Add("cmd", "EARSIV_PORTAL_MUSTAHSIL_GETIR")
		body.Add("pageName", "RG_MUSTAHSIL")
	case document.SelfEmployedReceipt:
		body.Add("cmd", "EARSIV_PORTAL_SERBEST_MESLEK_GETIR")
		body.Add("pageName", "RG_SERBEST")
	default:
		return nil, errors.New("invalid document type")
	}

	body.Add("ettn", id.String())

	response, err := gateway(DISPATCH, f.debug, body)
	if err != nil {
		return nil, err
	}

	if response.Has("error") {
		return nil, errors.New("document not found")
	}

	return response, nil
}

func (f *Fatura) GetDownloadURL(id uuid.UUID, signed bool) (string, error) {
	var body = types.Array{}

	body.Add("token", f.token)
	body.Add("ettn", id.String())
	body.Add("cmd", "EARSIV_PORTAL_FATURA_GOSTER")
	body.Add("pageName", "RG_TASLAKLAR")
	body.Add("onayDurumu", "Onaylanmadı")
	if signed {
		body.Add("onayDurumu", "Onaylandı")
	}

	response, err := gateway(DISPATCH, f.debug, body)
	if err != nil {
		return "", err
	}

	if response.Has("error") {
		return "", errors.New("download url not found")
	}

	if !response.Has("url") {
		return "", errors.New("download url not found")
	}

	return response.Get("url"), nil
}

func (f *Fatura) CancellationRequest(id uuid.UUID, explanation string) (string, error) {
	var body = types.Array{}

	body.Add("token", f.token)
	body.Add("ettn", id.String())
	body.Add("cmd", "EARSIV_PORTAL_FATURA_IPTAL_ISTEK")
	body.Add("pageName", "RG_TASLAKLAR")
	body.Add("belgeTuru", "FATURA")
	body.Add("talepAciklama", explanation)

	response, err := gateway(DISPATCH, f.debug, body)
	if err != nil {
		return "", errors.New("cancellation request failed")
	}

	if response.Has("error") {
		return "", errors.New("cancellation request failed")
	}

	if !response.Has("talepNo") {
		return "", errors.New("cancellation request failed")
	}

	return response.Get("talepNo"), nil
}

func (f *Fatura) GetHtml(id uuid.UUID, signed bool) ([]byte, error) {
	url, err := f.GetDownloadURL(id, signed)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var buf bytes.Buffer
	i, err := io.Copy(&buf, res.Body)
	if err != nil {
		return nil, err
	}

	if i == 0 {
		return nil, errors.New("download url not found")
	}

	return buf.Bytes(), nil
}

func (f *Fatura) GetPdf(id uuid.UUID, signed bool) (io.Reader, error) {
	html, err := f.GetHtml(id, signed)
	if err != nil {
		return nil, err
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.Write(0, string(html))

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}
