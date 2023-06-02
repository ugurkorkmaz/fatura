package document

import "github.com/ugurkorkmaz/fatura/enum"

type Type int

const (
	DocumentInvoice Type = iota + enum.HeadEnumDocument
	DocumentProducerReceipt
	DocumentSelfEmployedReceipt
)

func (t Type) String() string {
	return [...]string{
		"FATURA",
		"MÜSTAHSİL MAKBUZU",
		"SERBEST MESLEK MAKBUZU",
	}[t-enum.HeadEnumDocument]
}
