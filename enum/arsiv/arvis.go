package arsiv

import "github.com/ugurkorkmaz/fatura/enum"

// Type represents the enum constant type of the E-Arşiv.
type Type int

const (
	EArsiv Type = iota + enum.HeadEnumArsiv
	EArsivBuyuk
)

// String returns the string representation of the enum constant.
func (t Type) String() string {
	return [...]string{"EARSIV", "EARSIV_BUYUK"}[t-enum.HeadEnumArsiv]
}

// Alias returns the alias representation of the enum constant.
func (t Type) Alias() string {
	return [...]string{"E-Arşiv", "E-Arşiv Büyük"}[t-enum.HeadEnumArsiv]
}
