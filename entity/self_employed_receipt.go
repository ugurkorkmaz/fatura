package entity

import "encoding/json"

type SelfEmployedReceipt struct{}

func (s *SelfEmployedReceipt) Json() string {
	b, _ := json.Marshal(s)
	return string(b)
}
