package entity

import "encoding/json"

type User struct {
	VknTckn         string `json:"vknTckn"`
	Unvan           string `json:"unvan"`
	Ad              string `json:"ad"`
	Soyad           string `json:"soyad"`
	Cadde           string `json:"cadde"`
	ApartmanAdi     string `json:"apartmanAdi"`
	ApartmanNo      string `json:"apartmanNo"`
	KapiNo          string `json:"kapiNo"`
	Kasaba          string `json:"kasaba"`
	Ilce            string `json:"ilce"`
	Il              string `json:"il"`
	PostaKodu       string `json:"postaKodu"`
	Ulke            string `json:"ulke"`
	TelNo           string `json:"telNo"`
	FaksNo          string `json:"faksNo"`
	EPostaAdresi    string `json:"ePostaAdresi"`
	WebSitesiAdresi string `json:"webSitesiAdresi"`
	VergiDairesi    string `json:"vergiDairesi"`
	SicilNo         string `json:"sicilNo"`
	IsMerkezi       string `json:"isMerkezi"`
	MersisNo        string `json:"mersisNo"`
}

func (u *User) Name() string {
	return "KULLANICI"
}

func (u *User) Json() string {
	b, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(b)
}
