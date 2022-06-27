package dateFormatter

import (
	"time"
)

func GetStartOfDay(epoch int64) int64 {
	fecha := time.UnixMilli(epoch)
	year, month, day := fecha.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC).UnixMilli()
}

func GetEndOfDay(epoch int64) int64 {
	fecha := time.UnixMilli(epoch)
	year, month, day := fecha.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, time.UTC).UnixMilli()
}
