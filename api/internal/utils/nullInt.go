package utils

import "database/sql"

func NewNullInt64(value int64, nilDigit int64) sql.NullInt64 {
	if value != nilDigit {
		return sql.NullInt64{
			Int64: value,
			Valid: true,
		}
	}
	return sql.NullInt64{
		Int64: 0,
		Valid: false,
	}
}
