package entity

import (
	"github.com/ugurkorkmaz/fatura/enum/unit"
)

type ProducerItem struct {
	MalHizmet       string    `json:"malHizmet"`
	Miktar          float64   `json:"miktar"`
	Birim           unit.Type `json:"birim"`
	BirimFiyat      float64   `json:"birimFiyat"`
	MalHizmetTutari float64   `json:"malHizmetTutari"`
	GvStopajOrani   int       `json:"gvStopajOrani"`
}
