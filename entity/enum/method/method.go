package method

type Type int

const (
	Noter Type = iota
	TaahhutluMektup
	Telgraf
	Kep
)

func (t Type) String() string {
	return [...]string{
		"NOTER",
		"TAAHHUTLU_MEKTUP",
		"TELEFON",
		"KEP",
	}[t]
}
