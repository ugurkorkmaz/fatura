package entity

import (
	"encoding/json"
	"fatura/entity/enum/currency"

	"github.com/google/uuid"
)

type SelfEmployedReceipt struct {
	VknTckn           string           `json:"vknTckn"`
	AliciAdi          string           `json:"aliciAdi"`
	AliciSoyadi       string           `json:"aliciSoyadi"`
	Ulke              string           `json:"ulke"`
	Ettn              uuid.UUID        `json:"ettn"`
	BelgeNumarasi     string           `json:"belgeNumarasi"`
	Tarih             string           `json:"tarih"`
	Saat              string           `json:"saat"`
	ParaBirimi        currency.Type    `json:"paraBirimi"`
	DovizKuru         float64          `json:"dovizKuru"`
	AliciUnvan        string           `json:"aliciUnvan"`
	Adres             string           `json:"adres"`
	BinaAdi           string           `json:"binaAdi"`
	BinaNo            string           `json:"binaNo"`
	KapiNo            string           `json:"kapiNo"`
	KasabaKoy         string           `json:"kasabaKoy"`
	MahalleSemtIlce   string           `json:"mahalleSemtIlce"`
	Sehir             string           `json:"sehir"`
	PostaKodu         string           `json:"postaKodu"`
	VergiDairesi      string           `json:"vergiDairesi"`
	Aciklama          string           `json:"aciklama"`
	KdvTahakkukIcin   bool             `json:"kdvTahakkukIcin"`
	MalHizmetTable    SelfEmployedItem `json:"malHizmetTable"`
	BrutUcret         float64          `json:"brutUcret"`
	GvStopajTutari    float64          `json:"gvStopajTutari"`
	NetUcretTutari    float64          `json:"net"`
	KdvTutari         float64          `json:"kdvTutari"`
	KdvTevkifatTutari float64          `json:"kdv"`
	TahsilEdilenKdv   float64          `json:"tahsilEdilenKdv"`
	NetAlinanToplam   float64          `json:"netAlinanToplam"`
	Xxx               float64          `json:"xxx"`
}

func (s *SelfEmployedReceipt) Type() string {
	return "SERBEST MESLEK MAKBUZU"
}

func (s *SelfEmployedReceipt) Json() string {
	b, _ := json.Marshal(s)
	return string(b)
}

func (s *SelfEmployedReceipt) AddItem(item SelfEmployedItem) {}
