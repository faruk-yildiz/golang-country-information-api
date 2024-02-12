package dto

import "database/sql"

type UpdateOrAddCountryDto struct {
	Iso       sql.NullString `json:"iso"`
	Name      string         `json:"name"`
	NiceName  string         `json:"niceName"`
	Iso3      sql.NullString `json:"iso3"`
	NumCode   sql.NullInt32  `json:"numcode"`
	PhoneCode sql.NullInt32  `json:"phoneCode"`
}
