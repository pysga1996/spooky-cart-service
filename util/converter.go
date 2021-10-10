package util

import (
	"database/sql"
	"time"
)

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
