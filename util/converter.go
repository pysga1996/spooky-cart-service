package util

import (
	"database/sql"
	"time"
)

func GetNullableString(c *sql.NullString) *string {
	if c.Valid {
		return &c.String
	} else {
		return nil
	}
}

func GetNullableTime(c *sql.NullTime) *time.Time {
	if c.Valid {
		return &c.Time
	} else {
		return nil
	}
}

func GetNullableByte(c *sql.NullByte) *uint8 {
	if c.Valid {
		return &c.Byte
	} else {
		return nil
	}
}
