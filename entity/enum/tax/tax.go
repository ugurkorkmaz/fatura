package tax

type Type int

const (
	BankaMuameleleri Type = iota
	KKDFKesintisi
	OTV1Liste
	OTV2Liste
	OTV3Liste
	OTV4Liste
	OTV3AListe
	OTV3BListe
	OTV3CListe
	Damga
	Damga5035
	OzelIletisim
	OzelIletisim5035
	KDVTevkifat // Tevkifat
	BSMV4961
	BorsaTescil
	EnerjiFonu
	ElkHavagazTuketim
	TRTPayi
	ElkTuketim
	TKKullanim
	TKRuhsat
	CevreTemizlik
	GVStopaj          // Stopaj
	KVStopaj          // Stopaj
	MeraFonu          // Stopaj
	OTV1ListeTevkifat // Tevkifat
	BelOdHalRusum
	Konaklama
	caserim // Müstahsil
	SGKPrim
)

func (t Type) String() string {
	return [...]string{
		"0021", "0061", "0071", "9077", "0073", "0074", "0075",
		"0076", "0077", "1047", "1048", "4080", "4081", "9015",
		"9021", "8001", "8002", "4071", "8004", "8005", "8006",
		"8007", "8008", "0003", "0011", "9040", "4171", "9944",
		"0059", "SGK_PRIM",
	}[t]
}

func (t Type) Alias() string {
	return [...]string{
		"Banka Muameleleri Vergisi", "KKDF Kesintisi", "ÖTV 1. Liste", "ÖTV 2. Liste",
		"ÖTV 3. Liste", "ÖTV 4. Liste", "ÖTV 3A Liste", "ÖTV 3B Liste", "ÖTV 3C Liste",
		"Damga Vergisi", "5035 Sayılı Kanuna Göre Damga Vergisi", "Özel İletişim Vergisi",
		"5035 Sayılı Kanuna Göre Özel İletişim Vergisi", "KDV Tevkifat",
		"Banka ve Sigorta Muameleleri Vergisi", "Borsa Tescil Ücreti", "Enerji Fonu",
		"Elektrik Havagaz Tüketim Vergisi", "TRT Payı", "Elektrik Tüketim Vergisi",
		"TK Kullanım", "TK Ruhsat", "Çevre Temizlik Vergisi", "Gelir Vergisi Stopajı",
		"Kurumlar Vergisi Stopajı", "Mera Fonu", "ÖTV 1. Liste Tevkifat",
		"Belediyelere Ödenen Hal Rüsumu", "Konaklama Vergisi", "SGK Prim Kesintisi",
	}[t]
}

func (t Type) HasVat() bool {
	switch t {
	case
		KKDFKesintisi,
		OTV1Liste,
		OTV2Liste,
		OTV3Liste,
		OTV4Liste,
		OTV3AListe,
		OTV3BListe,
		OTV3CListe,
		EnerjiFonu,
		ElkHavagazTuketim,
		TRTPayi,
		ElkTuketim,
		OTV1ListeTevkifat,
		BelOdHalRusum:
		return true
	default:
		return false
	}
}

func (t Type) IsStoppage() bool {
	switch t {
	case
		KDVTevkifat,
		GVStopaj,
		KVStopaj,
		MeraFonu,
		SGKPrim:
		return true
	default:
		return false
	}
}

func (t Type) IsWithholding() bool {
	switch t {
	case
		KDVTevkifat,
		OTV1ListeTevkifat:
		return true
	default:
		return false
	}
}

func (t Type) HasDefaultRate() bool {
	switch t {
	case
		OTV1Liste,
		OTV1ListeTevkifat:
		return true
	default:
		return false
	}
}
func (t Type) DefaultRate() int {
	switch t {
	case OTV1Liste:
		return 0
	case OTV1ListeTevkifat:
		return 100
	}
	return -1
}
func (t Type) Codes(code any) (int, string) {
	switch code {
	case 601:
		return 40, "Yapım İşleri ile Bu İşlerle Birlikte İfa Edilen Mühendislik-Mimarlık ve Etüt-Proje Hizmetleri [KDVGUT-(I/C-2.1.3.2.1)]"
	case 602:
		return 90, "Etüt, plan-proje, danışmanlık, denetim vb"
	case 603:
		return 70, "Makine, Teçhizat, Demirbaş ve Taşıtlara Ait Tadil, Bakım ve Onarım Hizmetleri [KDVGUT- (I/C-2.1.3.2.3)]"
	case 604:
		return 50, "Yemek servis hizmeti"
	case 605:
		return 50, "Organizasyon hizmeti"
	case 606:
		return 90, "İşgücü temin hizmetleri"
	case 607:
		return 90, "Özel güvenlik hizmeti"
	case 608:
		return 90, "Yapı denetim hizmetleri"
	case 609:
		return 70, "Fason Olarak Yaptırılan Tekstil ve Konfeksiyon İşleri, Çanta ve Ayakkabı Dikim İşleri ve Bu İşlere Aracılık Hizmetleri [KDVGUT-(I/C-2.1.3.2.7)]"
	case 610:
		return 90, "Turistik mağazalara verilen müşteri bulma/ götürme hizmetleri"
	case 611:
		return 90, "Spor kulüplerinin yayın, reklam ve isim hakkı gelirlerine konu işlemleri"
	case 612:
		return 90, "Temizlik Hizmeti [KDVGUT-(I/C-2.1.3.2.10)]"
	case 613:
		return 90, "Çevre, Bahçe ve Bakım Hizmetleri [KDVGUT-(I/C-2.1.3.2.11)]"
	case 614:
		return 50, "Servis taşımacıliğı"
	case 615:
		return 70, "Her Türlü Baskı ve Basım Hizmetleri [KDVGUT-(I/C-2.1.3.2.12)]"
	case 616:
		return 50, "Diğer Hizmetler [KDVGUT-(I/C-2.1.3.2.13)]"
	case 617:
		return 70, "Hurda metalden elde edilen külçe teslimleri"
	case 618:
		return 70, "Hurda Metalden Elde Edilenler Dışındaki Bakır, Çinko, Demir Çelik, Alüminyum ve Kurşun Külçe Teslimi [KDVGUT-(I/C-2.1.3.3.1)]"
	case 619:
		return 70, "Bakir, çinko ve alüminyum ürünlerinin teslimi"
	case 620:
		return 70, "istisnadan vazgeçenlerin hurda ve atık teslimi"
	case 621:
		return 90, "Metal, plastik, lastik, kauçuk, kâğit ve cam hurda ve atıklardan elde edilen hammadde teslimi"
	case 622:
		return 90, "Pamuk, tiftik, yün ve yapaği ile ham post ve deri teslimleri"
	case 623:
		return 50, "Ağaç ve orman ürünleri teslimi"
	case 624:
		return 20, "Yük Taşımacılığı Hizmeti [KDVGUT-(I/C-2.1.3.2.11)]"
	case 625:
		return 30, "Ticari Reklam Hizmetleri [KDVGUT-(I/C-2.1.3.2.15)]"
	case 626:
		return 20, "Diğer Teslimler [KDVGUT-(I/C-2.1.3.3.7.)]"
	case 627:
		return 50, "Demir-Çelik Ürünlerinin Teslimi [KDVGUT-(I/C-2.1.3.3.8)]"
	case "627-Ex":
		return 40, "Demir-Çelik Ürünlerinin Teslimi [KDVGUT-(I/C-2.1.3.3.8)] (01/11/2022 tarihi öncesi)"
	case 801:
		return 100, "[Tam Tevkifat] Yapım İşleri ile Bu İşlerle Birlikte İfa Edilen Mühendislik-Mimarlık ve Etüt-Proje Hizmetleri[KDVGUT-(I/C-2.1.3.2.1)]"
	case 802:
		return 100, "[Tam Tevkifat] Etüt, Plan-Proje, Danışmanlık, Denetim ve Benzeri Hizmetler[KDVGUT-(I/C-2.1.3.2.2)]"
	case 803:
		return 100, "[Tam Tevkifat] Makine, Teçhizat, Demirbaş ve Taşıtlara Ait Tadil, Bakım ve Onarım Hizmetleri[KDVGUT- (I/C-2.1.3.2.3)]"
	case 804:
		return 100, "[Tam Tevkifat] Yemek Servis Hizmeti[KDVGUT-(I/C-2.1.3.2.4)]"
	case 805:
		return 100, "[Tam Tevkifat] Organizasyon Hizmeti[KDVGUT-(I/C-2.1.3.2.4)]"
	case 806:
		return 100, "[Tam Tevkifat] İşgücü Temin Hizmetleri[KDVGUT-(I/C-2.1.3.2.5)]"
	case 807:
		return 100, "[Tam Tevkifat] Özel Güvenlik Hizmeti[KDVGUT-(I/C-2.1.3.2.5)]"
	case 808:
		return 100, "[Tam Tevkifat] Yapı Denetim Hizmetleri[KDVGUT-(I/C-2.1.3.2.6)]"
	case 809:
		return 100, "[Tam Tevkifat] Fason Olarak Yaptırılan Tekstil ve Konfeksiyon İşleri, Çanta ve Ayakkabı Dikim İşleri ve Bu İşlere Aracılık Hizmetleri[KDVGUT-(I/C-2.1.3.2.7)]"
	case 810:
		return 100, "[Tam Tevkifat] Turistik Mağazalara Verilen Müşteri Bulma/ Götürme Hizmetleri[KDVGUT-(I/C-2.1.3.2.8)]"
	case 811:
		return 100, "[Tam Tevkifat] Spor Kulüplerinin Yayın, Reklâm ve İsim Hakkı Gelirlerine Konu İşlemleri[KDVGUT-(I/C-2.1.3.2.9)]"
	case 812:
		return 100, "[Tam Tevkifat] Temizlik Hizmeti[KDVGUT-(I/C-2.1.3.2.10)]"
	case 813:
		return 100, "[Tam Tevkifat] Çevreve Bahçe Bakım Hizmetleri[KDVGUT-(I/C-2.1.3.2.10)]"
	case 814:
		return 100, "[Tam Tevkifat] Servis Taşımacılığı Hizmeti[KDVGUT-(I/C-2.1.3.2.11)]"
	case 815:
		return 100, "[Tam Tevkifat] Her Türlü Baskı ve Basım Hizmetleri[KDVGUT-(I/C-2.1.3.2.12)]"
	case 816:
		return 100, "[Tam Tevkifat] Hurda Metalden Elde Edilen Külçe Teslimleri[KDVGUT-(I/C-2.1.3.3.1)]"
	case 817:
		return 100, "[Tam Tevkifat] Hurda Metalden Elde Edilenler Dışındaki Bakır, Çinko, Demir Çelik, Alüminyum ve Kurşun Külçe Teslimi [KDVGUT-(I/C-2.1.3.3.1)]"
	case 818:
		return 100, "[Tam Tevkifat] Bakır, Çinko, Alüminyum ve Kurşun Ürünlerinin Teslimi[KDVGUT-(I/C-2.1.3.3.2)]"
	case 819:
		return 100, "[Tam Tevkifat] İstisnadan Vazgeçenlerin Hurda ve Atık Teslimi[KDVGUT-(I/C-2.1.3.3.3)]"
	case 820:
		return 100, "[Tam Tevkifat] Metal, Plastik, Lastik, Kauçuk, Kâğıt ve Cam Hurda ve Atıklardan Elde Edilen Hammadde Teslimi[KDVGUT-(I/C-2.1.3.3.4)]"
	case 821:
		return 100, "[Tam Tevkifat] Pamuk, Tiftik, Yün ve Yapağı İle Ham Post ve Deri Teslimleri[KDVGUT-(I/C-2.1.3.3.5)]"
	case 822:
		return 100, "[Tam Tevkifat] Ağaç ve Orman Ürünleri Teslimi[KDVGUT-(I/C-2.1.3.3.6)]"
	case 823:
		return 100, "[Tam Tevkifat] Yük Taşımacılığı Hizmeti [KDVGUT-(I/C-2.1.3.2.11)]"
	case 824:
		return 100, "[Tam Tevkifat] Ticari Reklam Hizmetleri [KDVGUT-(I/C-2.1.3.2.15)]"
	case 825:
		return 100, "[Tam Tevkifat] Demir-Çelik Ürünlerinin Teslimi [KDVGUT-(I/C-2.1.3.3.8)]"
	}
	return -1, ""
}
