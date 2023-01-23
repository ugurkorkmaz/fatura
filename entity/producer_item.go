package entity

import "fatura/entity/enum/unit"

type ProducerItem struct {
	MalHizmet       string
	Miktar          float64
	Birim           unit.Type
	BirimFiyat      float64
	MalHizmetTutari float64
	GvStopajOrani   int
}
