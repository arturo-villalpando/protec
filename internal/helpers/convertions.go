package helpers

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func UUIDToString(uuid *pgtype.UUID) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid.Bytes[0:4], uuid.Bytes[4:6], uuid.Bytes[6:8], uuid.Bytes[8:10], uuid.Bytes[10:16])
}

func StringToUUID(src string) pgtype.UUID {
	data, err := ParseUUID(src)
	if err != nil {
		return pgtype.UUID{
			Bytes: [16]byte{},
			Valid: false,
		}
	}
	return pgtype.UUID{
		Bytes: data,
		Valid: true,
	}
}

func ParseUUID(src string) (dst [16]byte, err error) {
	switch len(src) {
	case 36:
		src = src[0:8] + src[9:13] + src[14:18] + src[19:23] + src[24:]
	case 32:
		// dashes already stripped, assume valid
	default:
		// assume invalid.
		return dst, fmt.Errorf("cannot parse UUID %v", src)
	}

	buf, err := hex.DecodeString(src)
	if err != nil {
		return dst, err
	}

	copy(dst[:], buf)
	return dst, err
}

func ParseTime(d string) *time.Time {
	if d != "" {
		date, _ := time.Parse(time.TimeOnly, d)

		return &date
	} else {
		return nil
	}
}

func ParseDate(d string) *time.Time {
	if d != "" {
		date, _ := time.Parse("2006-01-02", d)

		return &date
	} else {
		return nil
	}
}

func ParseDateTime(d string) *time.Time {
	if d != "" {
		date, _ := time.Parse("2006-01-02 15:04:05", d)

		return &date
	} else {
		return nil
	}
}

func ParseTimeStamp(d string) *time.Time {
	if d != "" {
		date, _ := time.Parse("2006-01-02T15:04:05Z", d)

		return &date
	} else {
		return nil
	}
}
