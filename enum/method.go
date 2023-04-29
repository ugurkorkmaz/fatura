package enum

type Method int

const (
	Noter Method = iota + headEnumMethod
	TaahhutluMektup
	Telgraf
	Kep
)

func (m Method) String() string {
	return [...]string{"NOTER", "TAAHHUTLU_MEKTUP", "TELEFON", "KEP"}[m-headEnumMethod]
}
