package entity

import (
	"encoding/json"

	"github.com/google/uuid"
)

type ProducerReceipt struct {
	Ettn                     uuid.UUID    `json:"ettn"`
	VknTckn                  string       `json:"vknTckn"`
	AliciAdi                 string       `json:"aliciAdi"`
	AliciSoyadi              string       `json:"aliciSoyadi"`
	BelgeNumarasi            string       `json:"belgeNumarasi"`
	Tarih                    string       `json:"tarih"`
	Saat                     string       `json:"saat"`
	Sehir                    string       `json:"sehir"`
	Websitesi                string       `json:"websitesi"`
	MalHizmetTable           ProducerItem `json:"malHizmetTable"`
	Not                      string       `json:"not"`
	TeslimTarihi             string       `json:"teslimTarihi"`
	MalHizmetToplamTutari    float64      `json:"malHizmetToplamTutari"`
	VergilerDahilToplamTutar float64      `json:"vergilerDahilToplamTutar"`
	OdenecekTutar            float64      `json:"odenecekTutar"`
}

func (p *ProducerReceipt) Type() string {
	return "MÜSTAHSİL MAKBUZU"
}
func (p *ProducerReceipt) Json() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}
