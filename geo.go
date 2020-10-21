package lalamove

// Country ...
type Country struct {
	Name       string
	Code       CountryCode
	Cities     []CityCode
	Locales    []Locale
	PhoneRegex string
}

// CityCode is the UN/LOCODE of supported cities.
type CityCode string

const (
	CityCodeBrasilSaoPaulo      CityCode = "BR_SAO"
	CityCodeBrasilRioDeJaneiro  CityCode = "BR_RIO"
	CityCodeHongKongHongKong    CityCode = "HK_HKG"
	CityCodeIndiaBengaluru      CityCode = "IN_BLR"
	CityCodeIndiaMumbai         CityCode = "IN_BOM"
	CityCodeIndiaDelhi          CityCode = "IN_DEL"
	CityCodeIndonesiaJakarata   CityCode = "ID_JKT"
	CityCodeMalaysiaKualaLumpur CityCode = "MY_KUL"
	CityCodeMexicoMexico        CityCode = "MX_MEX"
	CityCodePhilippinesManila   CityCode = "PH_MNL"
	CityCodePhilippinesCebu     CityCode = "PH_CEB"
	CityCodeSingaporeSingapore  CityCode = "SG_SIN"
	CityCodeTaiwanTaipei        CityCode = "TW_TPE"
	CityCodeThailandBangkok     CityCode = "TH_BKK"
	CityCodeThailandPattaya     CityCode = "TH_PYX"
	CityCodeVietnamHoChiMinh    CityCode = "VN_SGN"
	CityCodeVietnamHanoi        CityCode = "VN_HAN"
)

func (c CityCode) getLLMCountry() string {
	switch c {
	case CityCodeHongKongHongKong,
		CityCodeIndonesiaJakarata,
		CityCodeMalaysiaKualaLumpur,
		CityCodeSingaporeSingapore,
		CityCodeTaiwanTaipei:
		country := c.GetCountry()
		return string(country.Code)
	default:
		return string(c)
	}
}

// GetCountry ...
func (c CityCode) GetCountry() Country {
	switch c {
	case CityCodeBrasilSaoPaulo, CityCodeBrasilRioDeJaneiro:
		return CountryBrasil
	case CityCodeHongKongHongKong:
		return CountryHongKong
	case CityCodeIndiaBengaluru, CityCodeIndiaMumbai, CityCodeIndiaDelhi:
		return CountryIndia
	case CityCodeIndonesiaJakarata:
		return CountryIndonesia
	case CityCodeMalaysiaKualaLumpur:
		return CountryMalaysia
	case CityCodeMexicoMexico:
		return CountryMexico
	case CityCodePhilippinesManila, CityCodePhilippinesCebu:
		return CountryPhilippines
	case CityCodeSingaporeSingapore:
		return CountrySingapore
	case CityCodeTaiwanTaipei:
		return CountryTaiwan
	case CityCodeThailandBangkok, CityCodeThailandPattaya:
		return CountryThailand
	case CityCodeVietnamHoChiMinh, CityCodeVietnamHanoi:
		return CountryVietnam
	}
	return CountryUnknown
}

// // CityCode is the ISO 3166-1 alpha-2 of supported countries and regions.
type CountryCode string

const (
	CountryCodeBrasil      CountryCode = "BR"
	CountryCodeHongKong    CountryCode = "HK"
	CountryCodeIndia       CountryCode = "IN"
	CountryCodeIndonesia   CountryCode = "ID"
	CountryCodeMalaysia    CountryCode = "MY"
	CountryCodeMexico      CountryCode = "MX"
	CountryCodePhilippines CountryCode = "PH"
	CountryCodeSingapore   CountryCode = "SG"
	CountryCodeTaiwan      CountryCode = "TW"
	CountryCodeThailand    CountryCode = "TH"
	CountryCodeVietnam     CountryCode = "VN"
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
	CountryUnknown = Country{
		Name: "Unknown",
	}
	CountryBrasil = Country{
		Name:       "Brasil",
		Code:       CountryCodeBrasil,
		Cities:     []CityCode{CityCodeBrasilSaoPaulo, CityCodeBrasilRioDeJaneiro},
		PhoneRegex: "^[0-9]{2}[9]{1}[0-9]{8}$",
		Locales:    []Locale{LocaleBrasilEN, LocaleBrasilPT},
	}
	CountryHongKong = Country{
		Name:       "Hong Kong",
		Code:       CountryCodeBrasil,
		Cities:     []CityCode{CityCodeHongKongHongKong},
		PhoneRegex: "^((?!999)([2-9][0-9]{7}))$",
		Locales:    []Locale{LocaleHongKongEN, LocaleHongKongZH},
	}
	CountryIndia = Country{
		Name:       "India",
		Code:       CountryCodeIndia,
		Cities:     []CityCode{CityCodeIndiaBengaluru, CityCodeIndiaMumbai, CityCodeIndiaDelhi},
		PhoneRegex: "^([6-9][0-9]{9}|22[0-9]{8})$",
		Locales:    []Locale{LocaleIndiaEN, LocaleIndiaHI, LocaleIndiaKN, LocaleIndiaMR},
	}
	CountryIndonesia = Country{
		Name:       "Indonesia",
		Code:       CountryCodeIndonesia,
		Cities:     []CityCode{CityCodeIndonesiaJakarata},
		PhoneRegex: "^0(8\\d{8,11}|21\\d{7,8})$",
		Locales:    []Locale{LocaleIndonesiaEN, LocaleIndonesiaID},
	}
	CountryMalaysia = Country{
		Name:       "Malaysia",
		Code:       CountryCodeMalaysia,
		Cities:     []CityCode{CityCodeMalaysiaKualaLumpur},
		PhoneRegex: "^0(1[1,5]?\\d{8}|[4-7,9]\\d{7}|8[2-9]\\d{6}|3\\d{8})$",
		Locales:    []Locale{LocaleMalaysiaEN, LocaleMalaysiaMS},
	}
	CountryMexico = Country{
		Name:       "Mexico",
		Code:       CountryCodeMexico,
		Cities:     []CityCode{CityCodeMexicoMexico},
		PhoneRegex: "^([+]+52?)?(\\d{3}?){2}\\d{4}$",
		Locales:    []Locale{LocaleMexicoEN, LocaleMexicoMX},
	}
	CountryPhilippines = Country{
		Name:       "Philippines",
		Code:       CountryCodePhilippines,
		Cities:     []CityCode{CityCodePhilippinesManila, CityCodePhilippinesCebu},
		PhoneRegex: "^09[0-9]{9}$|^0?2[0-9]{7}$|^0?32[0-9]{7}$",
		Locales:    []Locale{LocalePhilippinesEN},
	}
	CountrySingapore = Country{
		Name:       "Singapore",
		Code:       CountryCodeSingapore,
		Cities:     []CityCode{CityCodeSingaporeSingapore},
		PhoneRegex: "^[689]{1}[0-9]{7}$",
		Locales:    []Locale{LocaleSingaporeEN},
	}
	CountryTaiwan = Country{
		Name:       "Taiwan",
		Code:       CountryCodeTaiwan,
		Cities:     []CityCode{CityCodeTaiwanTaipei},
		PhoneRegex: "^0([1-8]{1}[0-9]{7,8}|9[0-9]{8})$",
		Locales:    []Locale{LocaleTaiwanZH},
	}
	CountryThailand = Country{
		Name:       "Thailand",
		Code:       CountryCodeThailand,
		Cities:     []CityCode{CityCodeThailandBangkok, CityCodeThailandPattaya},
		PhoneRegex: "^(0[0-9]{8,9}|[0-9]{4})$",
		Locales:    []Locale{LocaleThailandEN, LocaleThailandTH},
	}
	CountryVietnam = Country{
		Name:       "Vietnam",
		Code:       CountryCodeVietnam,
		Cities:     []CityCode{CityCodeVietnamHoChiMinh, CityCodeVietnamHanoi},
		PhoneRegex: "^0?(2|[35789])[0-9]{8}$|^02[48][0-9]{8}$",
		Locales:    []Locale{LocaleVietnamEN, LocaleVietnamVI},
	}
)

// AllCountriesByISOCode list all the countries and regions supported by Lalamove by ISO 3166-1 alpha-2 code
var AllCountriesByISOCode = map[CountryCode]Country{
	CountryCodeBrasil:      CountryBrasil,
	CountryCodeHongKong:    CountryHongKong,
	CountryCodeIndia:       CountryIndia,
	CountryCodeIndonesia:   CountryIndonesia,
	CountryCodeMalaysia:    CountryMalaysia,
	CountryCodeMexico:      CountryMexico,
	CountryCodePhilippines: CountryPhilippines,
	CountryCodeSingapore:   CountrySingapore,
	CountryCodeTaiwan:      CountryTaiwan,
	CountryCodeThailand:    CountryThailand,
	CountryCodeVietnam:     CountryVietnam,
}
