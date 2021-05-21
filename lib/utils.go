package utils

import "time"

func TimeFormat(time time.Time) string {
	t := time.Format("2006-01-02 15:04:05")
	return t
}
