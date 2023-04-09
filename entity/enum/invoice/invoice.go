// Invoice types
package invoice

import "fatura/entity/enum"

type Type int

const (
	// sale
	Satis Type = iota + enum.HeadEnumInvoice
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

// typeToString is a map that maps each invoice type to its string representation
var typeToString = map[Type]string{
	Satis:        "SATIS",
	Iade:         "IADE",
	Tevkifat:     "TEVKIFAT",
	Istisna:      "ISTISNA",
	OzelMatrah:   "OZELMATRAH",
	IhracKayitli: "IHRACKAYITLI",
}

// typeToAlias is a map that maps each invoice type to its localized string representation
var typeToAlias = map[Type]string{
	Satis:        "Satış",
	Iade:         "İade",
	Tevkifat:     "Tevkifat",
	Istisna:      "İstisna",
	OzelMatrah:   "Özel Matrah",
	IhracKayitli: "İhraç Kayıtlı",
}

// typeToReasons is a map that maps each special tax base reason code to its description for OzelMatrah invoice types
var typeToReasons = map[Type]map[int]string{
	OzelMatrah: {
		801: "Milli Piyango, Spor Toto vb. Oyunlar",
		802: "At yarışları ve diğer müşterek bahis ve talih oyunları",
		803: "Profesyonel Sanatçıların Yer Aldığı Gösteriler, Konserler, Profesyonel Sporcuların Katıldığı Sportif Faaliyetler, Maçlar, Yarışlar ve Yarışmalar",
		804: "Gümrük Depolarında ve Müzayede Mahallerinde Yapılan Satışla",
		805: "Altından Mamül veya Altın İçeren Ziynet Eşyaları İle Sikke Altınların Teslimi",
		806: "Tütün Mamülleri",
		807: "Muzır Neşriyat Kapsamındaki  Gazete, Dergi vb. Periyodik Yayınlar",
		808: "Gümüşten Mamul veya Gümüş İçeren Ziynet Eşyaları ile Sikke Gümüşlerin Teslimi",
		809: "Belediyeler taraf. yap. şehiriçi yolcu taşımacılığında kullanılan biletlerin ve kartların bayiler tarafından satışı",
		810: "Ön Ödemeli Elektronik Haberleşme Hizmetleri",
		811: "TŞOF Tarafından Araç Plakaları ile Sürücü Kurslarında Kullanılan Bir Kısım Evrakın Teslimi",
		812: "KDV Uygulanmadan Alınan İkinci El Motorlu Kara Taşıtı veya Taşınmaz Teslimi",
	},
}

// String method returns the string representation of the invoice type
func (t Type) String() string {
	return typeToString[t]
}

// Alias method returns the localized string representation of the invoice type
func (t Type) Alias() string {
	return typeToAlias[t]
}

// Reasons method returns a map of special tax base reasons for OzelMatrah invoice types
func (t Type) Reasons() map[int]string {
	return typeToReasons[t]
}
