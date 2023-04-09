package currency

import "fatura/entity/enum"

type Type int

func (c Type) String() string {
	return [...]string{
		"AFN", "DZD", "ARS", "AWG", "AUD", "AZM", "BSD", "BHD", "THB", "PAB", "BBD",
		"BYR", "BZD", "BMD", "VEB", "BOB", "BGN", "BRL", "BND", "BIF", "CAD", "CVE",
		"KYD", "GHC", "XAF", "XPF", "CLP", "COP", "KMF", "NIO", "CRC", "CUP", "CYP",
		"CZK", "DKK", "GMD", "MKD", "AED", "DJF", "STD", "DOP", "VND", "AMD", "XCD",
		"EGP", "SVC", "ETB", "EUR", "FKP", "FJD", "HUF", "CDF", "GIP", "XAU", "HTG",
		"PYG", "GNF", "GYD", "HKD", "UAH", "ISK", "INR", "IQD", "IRR", "JMD", "JOD",
		"KES", "PGK", "LAK", "EEK", "HRK", "KWD", "MWK", "AOA", "MMK", "GEL", "LVL",
		"LBP", "ALL", "HNL", "SLL", "ROL", "LRD", "LYD", "SZL", "LTL", "LSL", "MGF",
		"MYR", "MTL", "TMM", "MUR", "MZM", "MXN", "MDL", "MAD", "NGN", "ERN", "NAD",
		"NPR", "ANG", "YUM", "ILS", "TWD", "NZD", "KPW", "NOK", "BTN", "PEN", "MRO",
		"TOP", "PKR", "XPD", "MOP", "UYU", "PHP", "XPT", "GBP", "BWP", "QAR", "GTQ",
		"ZAR", "OMR", "KHR", "MVR", "RON", "IDR", "RUB", "RWF", "SAR", "XDR", "SCR",
		"XAG", "SGD", "SKK", "SBD", "SOS", "LKR", "KGS", "TJS", "SHP", "SDD", "SRG",
		"SEK", "CHF", "SYP", "BDT", "WST", "TZS", "KZT", "SIT", "TTD", "MNT", "TND",
		"TRY", "UGX", "USD", "UZS", "VUV", "KRW", "YER", "JPY", "CNY", "ZWD", "PLN",
	}[c]
}

func (c Type) Alias() string {
	return [...]string{
		"Afghani", "Algerian Dinar", "Argentine Peso", "Aruban Guilder", "Australian Dollar",
		"Azerbaijanian Manat", "Bahamian Dollar", "Bahraini Dinar", "Baht", "Balboa", "Barbados Dollar",
		"Belarussian Ruble", "Belize Dollar", "Bermudian Dollar", "Bolivar", "Boliviano",
		"Bulgarian Lev", "Brazilian Real", "Brunei Dollar", "Burundi Franc", "Canadian Dollar",
		"Cape Verde Escudo", "Cayman Islands Dollar", "Cedi", "CFA Franc", "CFP Franc",
		"Chilean Peso", "Colombian Peso", "Comoro Franc", "Cordoba Oro",
		"Costa Rican Colon", "Cuban Peso", "Cyprus Pound", "Czech Koruna", "Danish Krone",
		"Dalasi", "Denar", "Dirham", "Djibouti Franc", "Dobra", "Dominican Peso", "Dong", "Dram",
		"East Carribean Dollar", "Egyptian Pound", "El Salvador Colon", "Ethopian Birr", "Euro",
		"Falkland Islands Pound", "Fiji Dollar", "Forint", "Franc Congolais", "Gibraltar Pound",
		"Gold", "Gourde", "Guarani", "Guinea Franc", "Guyana Dollar", "HKD", "Hryvnia", "Iceland Krona",
		"Indian Rupee", "Iraqi Dinar", "Iranian Rial", "Jamaican Dollar", "Jordanian Dinar", "Kenyan Shilling",
		"Kina", "Kip", "Kroon", "Kuna", "Kuwaiti Dinar", "Kwacha", "Kwanza", "Kyat", "Lari", "Latvian Lats",
		"Lebanese Pound", "Lek", "Lempira", "Leone", "Leu", "Liberian Dollar", "Libyan Dinar", "Lilangeni",
		"Lithuanian Litas", "Loti", "Malagasy Franc", "Malaysian Ringgit", "Maltese Lira", "Manat",
		"Mauritius Rupee", "Metical", "Mexican Peso", "Moldovan Leu", "Morrocan Dirham", "Naira",
		"Nakfa", "Namibia Dollar", "Nepalese Rupee", "Netherlands Antillian Guilder", "New Dinar",
		"New Israeli Sheqel", "New Taiwan Dollar", "New Zealand Dollar", "North Korean Won",
		"Norwegian Krone", "Ngultrum", "Nuevo Sol", "Ouguiya", "Pa&apos;anga", "Pakistan Rupee",
		"Palladium", "Pataca", "Peso Uruguayo", "Philippine Peso", "Platinum", "Pound Sterling",
		"Pula", "Qatari Rial", "Quetzal", "Rand", "Rial Omani", "Riel", "Rufiyaa", "Rumen Leyi", "Rupiah",
		"Russian Ruble", "Rwanda Franc", "Saudi Riyal", "SDR", "Seychelles Rupee", "Silver", "Singapore Dollar",
		"Slovak Koruna", "Solomon Islands Dollar", "Somali Shilling", "Sri Lanka Rupee", "Som", "Somoni",
		"St. Helena Pound", "Sudanese Dinar", "Suriname Guilder", "Swedish Krona", "Swiss Franc", "Syrian Pound",
		"Taka", "Tala", "Tanzanian Shilling", "Tenge", "Tolar", "Trinidad and Tobago Dollar", "Tugrik",
		"Tunisian Dinar", "Turkish Lira", "Uganda Shilling", "American Dollar", "Uzbekistan Sum", "Vatu",
		"Won", "Yemeni Rial", "Yen", "Yuan Renminbi", "Zimbabwe Dollar", "Zloty",
	}[c]
}

const (
	AFN Type = iota + enum.HeadEnumCurrency
	DZD
	ARS
	AWG
	AUD
	AZM
	BSD
	BHD
	THB
	PAB
	BBD
	BYR
	BZD
	BMD
	VEB
	BOB
	BGN
	BRL
	BND
	BIF
	CAD
	CVE
	KYD
	GHC
	XAF
	XPF
	CLP
	COP
	KMF
	NIO
	CRC
	CUP
	CYP
	CZK
	DKK
	GMD
	MKD
	AED
	DJF
	STD
	DOP
	VND
	AMD
	XCD
	EGP
	SVC
	ETB
	EUR
	FKP
	FJD
	HUF
	CDF
	GIP
	XAU
	HTG
	PYG
	GNF
	GYD
	HKD
	UAH
	ISK
	INR
	IQD
	IRR
	JMD
	JOD
	KES
	PGK
	LAK
	EEK
	HRK
	KWD
	MWK
	AOA
	MMK
	GEL
	LVL
	LBP
	ALL
	HNL
	SLL
	ROL
	LRD
	LYD
	SZL
	LTL
	LSL
	MGF
	MYR
	MTL
	TMM
	MUR
	MZM
	MXN
	MDL
	MAD
	NGN
	ERN
	NAD
	NPR
	ANG
	YUM
	ILS
	TWD
	NZD
	KPW
	NOK
	BTN
	PEN
	MRO
	TOP
	PKR
	XPD
	MOP
	UYU
	PHP
	XPT
	GBP
	BWP
	QAR
	GTQ
	ZAR
	OMR
	KHR
	MVR
	RON
	IDR
	RUB
	RWF
	SAR
	XDR
	SCR
	XAG
	SGD
	SKK
	SBD
	SOS
	LKR
	KGS
	TJS
	SHP
	SDD
	SRG
	SEK
	CHF
	SYP
	BDT
	WST
	TZS
	KZT
	SIT
	TTD
	MNT
	TND
	TRY
	UGX
	USD
	UZS
	VUV
	KRW
	YER
	JPY
	CNY
	ZWD
	PLN
)
