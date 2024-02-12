package response

type CountryResponseDto struct {
	Id        int64  `json:"id"`
	Iso       string `json:"iso"`
	Name      string `json:"name"`
	NiceName  string `json:"niceName"`
	Iso3      string `json:"iso3"`
	NumCode   int32  `json:"numcode"`
	PhoneCode int32  `json:"phoneCode"`
}
