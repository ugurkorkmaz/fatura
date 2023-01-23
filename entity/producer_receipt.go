package entity

import "encoding/json"

type ProducerReceipt struct{}

func (p *ProducerReceipt) Type() string {
	return "MÜSTAHSİL MAKBUZU"
}
func (p *ProducerReceipt) Json() string {
	b, _ := json.Marshal(p)
	return string(b)
}
