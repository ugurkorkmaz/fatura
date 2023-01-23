package entity

type SelfEmployedItem struct {
	NeIcinAlindigi   string  `json:"neIcinAlindigi"`
	BrutUcret        float64 `json:"brutUcret"`
	KdvOrani         int     `json:"kdvOrani"`
	GvStopajOrani    int     `json:"gvStopajOrani"`
	NetUcret         float64 `json:"netUcret"`
	KdvTevkifatOrani int     `json:"kdv"`
	NetAlinan        float64 `json:"netAlinan"`
}
