package common

import "time"

// ReverseMap reverses the key-value map to value-key
func ReverseMap(m map[string]string) map[string]string {
	n := make(map[string]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

// SecondsToMilliseconds converts the seconds into milliseconds
func SecondsToMilliseconds(seconds int64) int64 {
	return seconds * 1000
}

// GetCurrentTimeInUnixMilli returns the current time in Unix milliseconds,
// adjusted by the specified timedelta in milliseconds.
func GetCurrentTimeInUnixMilli(timedelta int64) int64 {
	now := time.Now()
	return now.UnixMilli() + timedelta
}
