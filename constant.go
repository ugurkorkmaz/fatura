package fatura

import "net/http"

type (
	// Api is the API of the server.
	Api string
	// Path is the path of the API.
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
