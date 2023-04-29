package enum

type Invoice int

const (
	// sale
	Satis Invoice = iota + headEnumInvoice
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
func (i Invoice) String() string {
	return [...]string{"SATIS", "IADE", "TEVKIFAT", "ISTISNA", "OZELMATRAH", "IHRACKAYITLI"}[i-headEnumInvoice]
}

// Alias method returns the localized string representation of the invoice type
func (i Invoice) Alias() string {
	return [...]string{"Satış", "İade", "Tevkifat", "İstisna", "Özel Matrah", "İhraç Kayıtlı"}[i-headEnumInvoice]
}
