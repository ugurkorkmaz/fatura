package arsiv

type Type int

const (
	EArsiv Type = iota
	EArsivBuyuk
)

func (t Type) String() string {
	return [...]string{
		"EARSIV",
		"EARSIV_BUYUK",
	}[t]
}

func (t Type) Alias() string {
	return [...]string{
		"E-Arşiv",
		"E-Arşiv Büyük",
	}[t]
}
