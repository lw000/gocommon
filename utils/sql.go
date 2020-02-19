package tyutils

import "database/sql"

func SqlNullBoolToBool(v sql.NullBool) bool {
	if v.Valid {
		return v.Bool
	}
	return false
}

func BoolToSqlNullBool(v int64) sql.NullBool {
	var d sql.NullBool
	err := d.Scan(v)
	if err != nil {

	}
	return d
}

func SqlNullFloat64ToInt64(v sql.NullFloat64) float64 {
	if v.Valid {
		return v.Float64
	}
	return 0
}

func Float64ToSqlNullFloat64(v int64) sql.NullFloat64 {
	var d sql.NullFloat64
	err := d.Scan(v)
	if err != nil {

	}
	return d
}

func SqlNullInt64ToInt64(v sql.NullInt64) int64 {
	if v.Valid {
		return v.Int64
	}
	return 0
}

func Int64ToSqlNullInt64(v int64) sql.NullInt64 {
	var d sql.NullInt64
	err := d.Scan(v)
	if err != nil {

	}
	return d
}

func SqlNullStringToString(v sql.NullString) string {
	if v.Valid {
		return v.String
	}
	return ""
}

func StringToSqlNullString(v string) sql.NullString {
	var d sql.NullString
	err := d.Scan(v)
	if err != nil {

	}
	return d
}
