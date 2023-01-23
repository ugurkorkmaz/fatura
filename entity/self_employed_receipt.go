package entity

import "encoding/json"

type SelfEmployedReceipt struct{}

func (s *SelfEmployedReceipt) Type() string {
	return "SERBEST MESLEK MAKBUZU"
}

func (s *SelfEmployedReceipt) Json() string {
	b, _ := json.Marshal(s)
	return string(b)
}
