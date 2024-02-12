package domain

import "database/sql"

type Country struct {
	Id        int64
	Iso       sql.NullString
	Name      string
	NiceName  string
	Iso3      sql.NullString
	NumCode   sql.NullInt32
	PhoneCode sql.NullInt32
}
