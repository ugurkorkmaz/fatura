package document

import "fatura/entity/enum"

// Document is a type of enum
type Type int

const (
	Invoice Type = iota + enum.HeadEnumDocument
	ProducerReceipt
	SelfEmployedReceipt
)

func (t Type) String() string {
	return [...]string{
		"FATURA",
		"MÜSTAHSİL MAKBUZU",
		"SERBEST MESLEK MAKBUZU",
	}[t]
}
