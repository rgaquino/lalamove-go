package lalamove

// Country ...
type Country struct {
	Name       string
	ISOCode    ISOCode
	UNLOCODE   []UNLOCODE
	PhoneRegex string
	LocaleKeys []Locale
}

// UNLOCODE is the UN/LOCODE of supported countries and regions.
type UNLOCODE string

const (
	UNLOCODEBrasilSaoPaulo      UNLOCODE = "BR_SAO"
	UNLOCODEBrasilRioDeJaneiro  UNLOCODE = "BR_RIO"
	UNLOCODEHongKongHongKong    UNLOCODE = "HK_HKG"
	UNLOCODEIndiaBengaluru      UNLOCODE = "IN_BLR"
	UNLOCODEIndiaMumbai         UNLOCODE = "IN_BOM"
	UNLOCODEIndiaDelhi          UNLOCODE = "IN_DEL"
	UNLOCODEIndonesiaJakarata   UNLOCODE = "ID_JKT"
	UNLOCODEMalaysiaKualaLumpur UNLOCODE = "MY_KUL"
	UNLOCODEMexicoMexico        UNLOCODE = "MX_MEX"
	UNLOCODEPhilippinesManila   UNLOCODE = "PH_MNL"
	UNLOCODEPhilippinesCebu     UNLOCODE = "PH_CEB"
	UNLOCODESingaporeSingapore  UNLOCODE = "SG_SIN"
	UNLOCODETaiwanTaipei        UNLOCODE = "TW_TPE"
	UNLOCODEThailandBangkok     UNLOCODE = "TH_BKK"
	UNLOCODEThailandPattaya     UNLOCODE = "TH_PYX"
	UNLOCODEVietnamHoChiMinh    UNLOCODE = "VN_SGN"
	UNLOCODEVietnamHanoi        UNLOCODE = "VN_HAN"
)

func (c UNLOCODE) getLLMCountry() string {
	switch c {
	case UNLOCODEHongKongHongKong,
		UNLOCODEIndonesiaJakarata,
		UNLOCODEMalaysiaKualaLumpur,
		UNLOCODESingaporeSingapore,
		UNLOCODETaiwanTaipei:
		country := AllCountriesByUNLOCODE[c]
		return string(country.ISOCode)
	default:
		return string(c)
	}
}

// // UNLOCODE is the ISO 3166-1 alpha-2 of supported countries and regions.
type ISOCode string

const (
	ISOCodeBrasil      ISOCode = "BR"
	ISOCodeHongKong    ISOCode = "HK"
	ISOCodeIndia       ISOCode = "IN"
	ISOCodeIndonesia   ISOCode = "ID"
	ISOCodeMalaysia    ISOCode = "MY"
	ISOCodeMexico      ISOCode = "MX"
	ISOCodePhilippines ISOCode = "PH"
	ISOCodeSingapore   ISOCode = "SG"
	ISOCodeTaiwan      ISOCode = "TW"
	ISOCodeThailand    ISOCode = "TH"
	ISOCodeVietnam     ISOCode = "VN"
)

// Locale ...
type Locale string

// Locale enum
const (
	LocaleBrasilEN      Locale = "en_BR"
	LocaleBrasilPT      Locale = "pt_BR"
	LocaleHongKongEN    Locale = "en_HK"
	LocaleHongKongZH    Locale = "zh_HK"
	LocaleIndiaEN       Locale = "en_IN"
	LocaleIndiaHI       Locale = "hi_IN"
	LocaleIndiaKN       Locale = "kn_IN"
	LocaleIndiaMR       Locale = "mr_IN"
	LocaleIndonesiaEN   Locale = "en_ID"
	LocaleIndonesiaID   Locale = "id_ID"
	LocaleMalaysiaEN    Locale = "en_MY"
	LocaleMalaysiaMS    Locale = "ms_MY"
	LocaleMexicoEN      Locale = "en_MX"
	LocaleMexicoMX      Locale = "es_MX"
	LocalePhilippinesEN Locale = "en_PH"
	LocaleSingaporeEN   Locale = "en_SG"
	LocaleTaiwanZH      Locale = "zh_TW"
	LocaleThailandEN    Locale = "en_TH"
	LocaleThailandTH    Locale = "th_TH"
	LocaleVietnamEN     Locale = "en_VN"
	LocaleVietnamVI     Locale = "vi_VN"
)

// Country enum
var (
	CountryBrasil = Country{
		Name:       "Brasil",
		ISOCode:    ISOCodeBrasil,
		UNLOCODE:   []UNLOCODE{UNLOCODEBrasilSaoPaulo, UNLOCODEBrasilRioDeJaneiro},
		PhoneRegex: "^[0-9]{2}[9]{1}[0-9]{8}$",
	}
	CountryHongKong = Country{
		Name:       "Hong Kong",
		ISOCode:    ISOCodeBrasil,
		UNLOCODE:   []UNLOCODE{UNLOCODEHongKongHongKong},
		PhoneRegex: "^((?!999)([2-9][0-9]{7}))$",
	}
	CountryIndia = Country{
		Name:       "India",
		ISOCode:    ISOCodeIndia,
		UNLOCODE:   []UNLOCODE{UNLOCODEIndiaBengaluru, UNLOCODEIndiaMumbai, UNLOCODEIndiaDelhi},
		PhoneRegex: "^([6-9][0-9]{9}|22[0-9]{8})$",
	}
	CountryIndonesia = Country{
		Name:       "Indonesia",
		ISOCode:    ISOCodeIndonesia,
		UNLOCODE:   []UNLOCODE{UNLOCODEIndonesiaJakarata},
		PhoneRegex: "^0(8\\d{8,11}|21\\d{7,8})$",
	}
	CountryMalaysia = Country{
		Name:       "Malaysia",
		ISOCode:    ISOCodeMalaysia,
		UNLOCODE:   []UNLOCODE{UNLOCODEMalaysiaKualaLumpur},
		PhoneRegex: "^0(1[1,5]?\\d{8}|[4-7,9]\\d{7}|8[2-9]\\d{6}|3\\d{8})$",
	}
	CountryMexico = Country{
		Name:       "Mexico",
		ISOCode:    ISOCodeMexico,
		UNLOCODE:   []UNLOCODE{UNLOCODEMexicoMexico},
		PhoneRegex: "^([+]+52?)?(\\d{3}?){2}\\d{4}$",
	}
	CountryPhilippines = Country{
		Name:       "Philippines",
		ISOCode:    ISOCodePhilippines,
		UNLOCODE:   []UNLOCODE{UNLOCODEPhilippinesManila, UNLOCODEPhilippinesCebu},
		PhoneRegex: "^09[0-9]{9}$|^0?2[0-9]{7}$|^0?32[0-9]{7}$",
	}
	CountrySingapore = Country{
		Name:       "Singapore",
		ISOCode:    ISOCodeSingapore,
		UNLOCODE:   []UNLOCODE{UNLOCODESingaporeSingapore},
		PhoneRegex: "^[689]{1}[0-9]{7}$",
	}
	CountryTaiwan = Country{
		Name:       "Taiwan",
		ISOCode:    ISOCodeTaiwan,
		UNLOCODE:   []UNLOCODE{UNLOCODETaiwanTaipei},
		PhoneRegex: "^0([1-8]{1}[0-9]{7,8}|9[0-9]{8})$",
	}
	CountryThailand = Country{
		Name:       "Thailand",
		ISOCode:    ISOCodeThailand,
		UNLOCODE:   []UNLOCODE{UNLOCODEThailandBangkok, UNLOCODEThailandPattaya},
		PhoneRegex: "^(0[0-9]{8,9}|[0-9]{4})$",
	}
	CountryVietnam = Country{
		Name:       "Vietnam",
		ISOCode:    ISOCodeVietnam,
		UNLOCODE:   []UNLOCODE{UNLOCODEVietnamHoChiMinh, UNLOCODEVietnamHanoi},
		PhoneRegex: "^0?(2|[35789])[0-9]{8}$|^02[48][0-9]{8}$",
	}
)

// AllCountriesByISOCode list all the countries and regions supported by Lalamove by ISO 3166-1 alpha-2 code
var AllCountriesByISOCode = map[ISOCode]Country{
	ISOCodeBrasil:      CountryBrasil,
	ISOCodeHongKong:    CountryHongKong,
	ISOCodeIndia:       CountryIndia,
	ISOCodeIndonesia:   CountryIndonesia,
	ISOCodeMalaysia:    CountryMalaysia,
	ISOCodeMexico:      CountryMexico,
	ISOCodePhilippines: CountryPhilippines,
	ISOCodeSingapore:   CountrySingapore,
	ISOCodeTaiwan:      CountryTaiwan,
	ISOCodeThailand:    CountryThailand,
	ISOCodeVietnam:     CountryVietnam,
}

// AllCountriesByUNLOCODE list all the countries and regions supported by Lalamove by UN/LOCODE
var AllCountriesByUNLOCODE = map[UNLOCODE]Country{
	UNLOCODEBrasilSaoPaulo:      CountryBrasil,
	UNLOCODEBrasilRioDeJaneiro:  CountryBrasil,
	UNLOCODEHongKongHongKong:    CountryHongKong,
	UNLOCODEIndiaBengaluru:      CountryIndia,
	UNLOCODEIndiaMumbai:         CountryIndia,
	UNLOCODEIndiaDelhi:          CountryIndia,
	UNLOCODEIndonesiaJakarata:   CountryIndonesia,
	UNLOCODEMalaysiaKualaLumpur: CountryMalaysia,
	UNLOCODEMexicoMexico:        CountryMexico,
	UNLOCODEPhilippinesManila:   CountryPhilippines,
	UNLOCODEPhilippinesCebu:     CountryPhilippines,
	UNLOCODESingaporeSingapore:  CountrySingapore,
	UNLOCODETaiwanTaipei:        CountryTaiwan,
	UNLOCODEThailandBangkok:     CountryThailand,
	UNLOCODEThailandPattaya:     CountryThailand,
	UNLOCODEVietnamHoChiMinh:    CountryVietnam,
	UNLOCODEVietnamHanoi:        CountryVietnam,
}
