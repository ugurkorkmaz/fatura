package entity

import (
	"encoding/json"
	"fatura/entity/enum/arsiv"
	"fatura/entity/enum/currency"
	"fatura/entity/enum/invoice"
)

type Array map[string]interface{}

type Invoice struct {
	trait `json:"-"`

	VknTckn                  string            `json:"vknTckn"`
	AliciAdi                 string            `json:"aliciAdi"`
	AliciSoyadi              string            `json:"aliciSoyadi"`
	MahalleSemtIlce          string            `json:"mahalleSemtIlce"`
	Sehir                    string            `json:"sehir"`
	Ulke                     string            `json:"ulke"`
	HangiTip                 arsiv.Type        `json:"hangiTip"`
	BelgeNumarasi            string            `json:"belgeNumarasi"`
	Tarih                    string            `json:"tarih"`
	Saat                     string            `json:"saat"`
	ParaBirimi               currency.Type     `json:"paraBirimi"`
	DovizKuru                float64           `json:"dovizKuru"`
	FaturaTipi               invoice.Type      `json:"faturaTipi"`
	SiparisNumarasi          string            `json:"siparisNumarasi"`
	SiparisTarihi            string            `json:"siparisTarihi"`
	IrsaliyeNumarasi         string            `json:"irsaliyeNumarasi"`
	IrsaliyeTarihi           string            `json:"irsaliyeTarihi"`
	FisNo                    string            `json:"fisNo"`
	FisTarihi                string            `json:"fisTarihi"`
	FisSaati                 string            `json:"fisSaati"`
	FisTipi                  string            `json:"fisTipi"`
	ZRaporNo                 string            `json:"zRaporNo"`
	OkcSeriNo                string            `json:"okcSeriNo"`
	AliciUnvan               string            `json:"aliciUnvan"`
	Adres                    string            `json:"adres"`
	BinaAdi                  string            `json:"binaAdi"`
	BinaNo                   string            `json:"binaNo"`
	KapiNo                   string            `json:"kapiNo"`
	KasabaKoy                string            `json:"kasabaKoy"`
	PostaKodu                string            `json:"postaKodu"`
	Tel                      string            `json:"tel"`
	Fax                      string            `json:"fax"`
	Eposta                   string            `json:"eposta"`
	Websitesi                string            `json:"websitesi"`
	VergiDairesi             string            `json:"vergiDairesi"`
	IadeTable                InvoiceReturnItem `json:"iadeTable"`
	MalHizmetTable           InvoiceItem       `json:"malHizmetTable"`
	Not                      string            `json:"not"`
	Matrah                   float64           `json:"matrah"`
	MalHizmetToplamTutari    float64           `json:"malHizmetToplamTutari"`
	ToplamIskonto            float64           `json:"toplamIskonto"`
	HesaplananKdv            float64           `json:"hesaplananKdv"`
	VergilerToplami          float64           `json:"vergilerToplami"`
	VergilerDahilToplamTutar float64           `json:"vergilerDahilToplamTutar"`
	ToplamMasraflar          float64           `json:"toplamMasraflar"`
	OdenecekTutar            float64           `json:"odenecekTutar"`
}

func (i *Invoice) Type() string {
	return "FATURA"
}

func (i *Invoice) Default() {}

func (i *Invoice) Validate() bool {
	return true
}

func (i *Invoice) Json() string {
	_json, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(_json)
}
