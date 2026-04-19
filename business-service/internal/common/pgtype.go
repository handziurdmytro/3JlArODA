package common

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func TextFromString(value string) pgtype.Text {
	return pgtype.Text{
		String: value,
		Valid:  true,
	}
}

func TextFromPtr(value *string) pgtype.Text {
	if value == nil {
		return pgtype.Text{}
	}

	return TextFromString(*value)
}

func PtrFromText(value pgtype.Text) *string {
	if !value.Valid {
		return nil
	}

	return &value.String
}

func DateFromTime(value time.Time) pgtype.Date {
	return pgtype.Date{
		Time:  value,
		Valid: true,
	}
}

func TimeFromDate(value pgtype.Date) time.Time {
	if !value.Valid {
		return time.Time{}
	}

	return value.Time
}

func TimestampFromTime(value time.Time) pgtype.Timestamp {
	return pgtype.Timestamp{
		Time:  value,
		Valid: true,
	}
}

func TimeFromTimestamp(value pgtype.Timestamp) time.Time {
	if !value.Valid {
		return time.Time{}
	}

	return value.Time
}

func NumericFromFloat64(value float64) (pgtype.Numeric, error) {
	var numeric pgtype.Numeric
	if err := numeric.Scan(strconv.FormatFloat(value, 'f', -1, 64)); err != nil {
		return pgtype.Numeric{}, fmt.Errorf("convert float64 to numeric: %w", err)
	}

	return numeric, nil
}

func Float64FromNumeric(value pgtype.Numeric) (float64, error) {
	floatValue, err := value.Float64Value()
	if err != nil {
		return 0, fmt.Errorf("convert numeric to float64: %w", err)
	}
	if !floatValue.Valid {
		return 0, nil
	}

	return floatValue.Float64, nil
}
