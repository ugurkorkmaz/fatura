package enum

type Arsiv int

const (
	EArsiv Arsiv = iota + headEnumArsiv
	EArsivBuyuk
)

// String returns the string representation of the enum constant.
func (a Arsiv) String() string {
	return [...]string{"EARSIV", "EARSIV_BUYUK"}[a-headEnumArsiv]
}

// Alias returns the alias representation of the enum constant.
func (a Arsiv) Alias() string {
	return [...]string{"E-Arşiv", "E-Arşiv Büyük"}[a-headEnumArsiv]
}
