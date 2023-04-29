package entity

import "encoding/json"

type SelfEmployedItem struct {
	NeIcinAlindigi   string  `json:"neIcinAlindigi"`
	BrutUcret        float64 `json:"brutUcret"`
	KdvOrani         int     `json:"kdvOrani"`
	GvStopajOrani    int     `json:"gvStopajOrani"`
	NetUcret         float64 `json:"netUcret"`
	KdvTevkifatOrani int     `json:"kdv"`
	NetAlinan        float64 `json:"netAlinan"`
}

func (s *SelfEmployedItem) Name() string {
	return "SERBEST MESLEK MAKBUZU MAL/HİZMET TABLOSU"
}

func (s *SelfEmployedItem) Json() (string, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
