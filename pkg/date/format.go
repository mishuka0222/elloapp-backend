package date

import "time"

var (
	ISO8601_WITH_TIMEZONE = "2006-01-02T15:04:05+0000"
)

func FormatDateIso8601(value string) (time.Time, error) {
	return time.Parse(ISO8601_WITH_TIMEZONE, value)
}
