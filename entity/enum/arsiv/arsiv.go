package arsiv

import "fatura/entity/enum"

type Type int

const (
	EArsiv Type = iota + enum.HeadEnumArsiv
	EArsivBuyuk
)

// typeToString is a map that holds the relationship between enum constants and their string values.
var typeToString = map[Type]string{
	EArsiv:      "EARSIV",
	EArsivBuyuk: "EARSIV_BUYUK",
}

// typeToAlias is a map that holds the relationship between enum constants and their alias values.
var typeToAlias = map[Type]string{
	EArsiv:      "E-Arşiv",
	EArsivBuyuk: "E-Arşiv Büyük",
}

// String returns the string representation of the enum constant.
// It looks up the value in the typeToString map.
func (t Type) String() string {
	return typeToString[t]
}

// Alias returns the alias representation of the enum constant.
// It looks up the value in the typeToAlias map.
func (t Type) Alias() string {
	return typeToAlias[t]
}
