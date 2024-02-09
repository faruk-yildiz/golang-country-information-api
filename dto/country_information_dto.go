package dto

type CountryDto struct {
	Iso       string `json:"iso"`
	Name      string `json:"name"`
	NiceName  string `json:"niceName"`
	Iso3      string `json:"iso3"`
	Numcode   int    `json:"numcode"`
	PhoneCode int    `json:"phoneCode"`
}
