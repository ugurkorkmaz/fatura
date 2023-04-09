package unit

import "fatura/entity/enum"

type Type int

const (
	Gun Type = iota + enum.HeadEnumUnit
	Ay
	Yil
	Saat
	Dk
	Sn
	Adet
	Pk
	Kutu
	Mgm
	Grm
	Kgm
	Ltr
	Ton
	Nt
	Gt
	Mmt
	Cmt
	Mtr
	Ktm
	Mlt
	Mm3
	Cm2
	Cmq
	M2
	M3
	Kjo
	Clt
	Ct
	Kwh
	Mwh
	Cct
	Gkj
	Klt
	Lpa
	Kgm2
	Ncl
	Pr
	Kmt
	Set
	T3
	Scm
	Ncm
	Mmbtu
	Cm3
	Dzn
	Dm2
	Dmt
	Har
	Lm
)

var (
	// unitTypeToString is a map of unit types to their string representation.
	unitTypeToString = [...]string{
		"DAY", "MON", "ANN", "HUR", "D61", "D62", "C62", "PA", "BX", "MGM", "GRM", "KGM",
		"LTR", "TNE", "NT", "GT", "MMT", "CMT", "MTR", "KTM", "MLT", "MMQ", "CMK", "CMQ",
		"MTK", "MTQ", "KJO", "CLT", "Ct", "KWH", "MWH", "CCT", "D30", "D40", "LPA", "B32",
		"NCL", "Pr", "R9", "SET", "T3", "Q37", "Q39", "J39", "G52", "DZN", "DMK", "DMT",
		"HAR", "LM",
	}
	// unitTypeToAlias is a map of unit types to their alias representation.
	unitTypeToAlias = [...]string{
		"Gün", "Ay", "Yıl", "Saat", "Dakika", "Saniye", "Adet", "Paket", "Kutu", "Mg", "Gram", "Kg", "Lt",
		"Ton", "Net Ton", "Gross ton", "Mm", "Cm", "M", "Km", "Ml", "Mm3", "Cm2", "Cm3", "M2", "M3", "Kj",
		"Cl", "Karat", "Kwh", "Mwh", "Ton Başına Taşıma Kapasitesi", "Brüt Kalori", "1000 Lt",
		"Saf Alkol Lt", "Kg M2", "Hücre Adet", "Çift", "1000 M3", "Set", "1000 Adet", "Scm", "Ncm",
		"Mmbtu", "Cm³", "Düzine", "Dm2", "Dm", "Ha", "Metretül (LM)",
	}
)

func (t Type) String() string {
	return unitTypeToString[t]
}

func (t Type) Alias() string {
	return unitTypeToAlias[t]
}
