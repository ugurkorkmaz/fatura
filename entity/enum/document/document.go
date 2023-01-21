package document

type Type int

const (
	Invoice Type = iota
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
