package db

import "database/sql"

//ToNullString returns a null string (otherwise it would be an empty string).
func ToNullString(s string) sql.NullString {
	obj := sql.NullString{}
	_ = obj.Scan(s)
	return obj
}
