package entity

import (
	"encoding/json"
	"fatura/entity/enum/unit"
)

type InvoiceItem struct {
	MalHizmet        string    `json:"malHizmet"`
	Miktar           float64   `json:"miktar"`
	Birim            unit.Type `json:"birim"`
	BirimFiyat       float64   `json:"birimFiyat"`
	KdvOrani         int       `json:"kdvOrani"`
	Fiyat            float64   `json:"fiyat"`
	IskontoTipi      bool      `json:"iskontoArttm"`
	IskontoOrani     float64   `json:"iskontoOrani"`
	IskontoTutari    float64   `json:"iskontoTutari"`
	IskontoNedeni    string    `json:"iskontoNedeni"`
	MalHizmetTutari  float64   `json:"malHizmetTutari"`
	KdvTutari        float64   `json:"kdvTutari"`
	TevkifatKodu     int       `json:"tevkifatKodu"`
	OzelMatrahNedeni int       `json:"ozelMatrahNedeni"`
	OzelMatrahTutari float64   `json:"ozelMatrahTutari"`
	Gtip             string    `json:"gtip"`
}

func (i *InvoiceItem) New() *InvoiceItem {
	return &InvoiceItem{
		Fiyat:            0,
		IskontoTipi:      false,
		IskontoOrani:     0,
		IskontoTutari:    0,
		IskontoNedeni:    "",
		MalHizmetTutari:  0,
		KdvTutari:        0,
		TevkifatKodu:     0,
		OzelMatrahNedeni: 0,
		OzelMatrahTutari: 0,
		Gtip:             "",
	}
}

// TODO
func (i *InvoiceItem) Default() {
	panic("implement me")
}

// TODO
func (i *InvoiceItem) Validate() bool {
	panic("implement me")
}

func (i *InvoiceItem) Json() string {
	_json, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(_json)
}
