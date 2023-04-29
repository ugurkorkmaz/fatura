package enum

type Document int

const (
	DocumentInvoice Document = iota + headEnumDocument
	DocumentProducerReceipt
	DocumentSelfEmployedReceipt
)

func (t Document) String() string {
	return [...]string{
		"FATURA",
		"MÜSTAHSİL MAKBUZU",
		"SERBEST MESLEK MAKBUZU",
	}[t-headEnumDocument]
}
