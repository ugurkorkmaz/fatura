package entity

import (
	"github.com/ugurkorkmaz/fatura/enum"
)

type ProducerItem struct {
	MalHizmet       string    `json:"malHizmet"`
	Miktar          float64   `json:"miktar"`
	Birim           enum.Unit `json:"birim"`
	BirimFiyat      float64   `json:"birimFiyat"`
	MalHizmetTutari float64   `json:"malHizmetTutari"`
	GvStopajOrani   int       `json:"gvStopajOrani"`
}
