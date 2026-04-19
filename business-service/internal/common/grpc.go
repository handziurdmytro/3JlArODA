package common

import (
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	DateLayout = "2006-01-02"
)

func ParseProtoTime(value string) (time.Time, error) {
	if value == "" {
		return time.Time{}, status.Error(codes.InvalidArgument, "date value is required")
	}

	parsed, err := time.Parse(time.RFC3339, value)
	if err == nil {
		return parsed, nil
	}

	parsed, err = time.Parse(DateLayout, value)
	if err == nil {
		return parsed, nil
	}

	return time.Time{}, status.Error(
		codes.InvalidArgument,
		fmt.Sprintf("invalid date %q: use RFC3339 or YYYY-MM-DD", value),
	)
}

func FormatProtoTime(value time.Time) string {
	if value.IsZero() {
		return ""
	}

	return value.Format(time.RFC3339)
}

func ToStatusError(err error) error {
	switch {
	case errors.Is(err, ErrNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, ErrAlreadyExists):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, ErrForeignKeyViolation),
		errors.Is(err, ErrCheckViolation),
		errors.Is(err, ErrNotNullViolation):
		return status.Error(codes.InvalidArgument, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}
