package invoice

import "github.com/ugurkorkmaz/fatura/enum"

type Type int

const (
	// sale
	Satis Type = iota + enum.HeadEnumInvoice
	// return
	Iade
	// tax withholding
	Tevkifat
	// exemption
	Istisna
	// special tax base
	OzelMatrah
	// export registered
	IhracKayitli
)

// String method returns the string representation of the invoice type
func (t Type) String() string {
	return [...]string{"SATIS", "IADE", "TEVKIFAT", "ISTISNA", "OZELMATRAH", "IHRACKAYITLI"}[t-enum.HeadEnumInvoice]
}

// Alias method returns the localized string representation of the invoice type
func (t Type) Alias() string {
	return [...]string{"Satış", "İade", "Tevkifat", "İstisna", "Özel Matrah", "İhraç Kayıtlı"}[t-enum.HeadEnumInvoice]
}
