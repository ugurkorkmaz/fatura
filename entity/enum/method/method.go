package method

import "fatura/entity/enum"

type Type int

const (
	Noter Type = iota + enum.HeadEnumMethod
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
