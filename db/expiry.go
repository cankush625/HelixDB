package db

import (
	"time"
)

const (
	cleanupInterval       = time.Second
	chunkSize             = 20
	expiredRatioThreshold = 0.25
)

// StartActiveExpiry launches a background goroutine that periodically scans
// KeyTTL for expired keys and removes them from the store.
// It processes keys in chunks of chunkSize to avoid long pauses.
// If the expired ratio within a chunk exceeds expiredRatioThreshold, it
// immediately runs another round rather than waiting for the next tick —
// this handles bursts of expiring keys efficiently.
func StartActiveExpiry() {
	go func() {
		ticker := time.NewTicker(cleanupInterval)
		defer ticker.Stop()
		for range ticker.C {
			for {
				total, expired := cleanExpiredKeys()
				if total < chunkSize || float64(expired)/float64(total) < expiredRatioThreshold {
					break
				}
			}
		}
	}()
}

// cleanExpiredKeys scans up to chunkSize keys from KeyTTL, deletes any that
// have expired, and returns the number of keys scanned and the number deleted.
func cleanExpiredKeys() (total int, expired int) {
	now := time.Now().UnixMilli()
	KeyTTL.Range(func(key, value any) bool {
		total++
		if value != nil && value.(int64) < now {
			KeyTTL.Delete(key)
			DB.Delete(key)
			expired++
		}
		return total < chunkSize
	})
	return
}
