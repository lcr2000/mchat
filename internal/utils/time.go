package utils

import "time"

func TimeFormat(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}
