package db

import (
	"time"
)

const (
	chunkSize             = 20
	expiredRatioThreshold = 0.25
)

// StartActiveExpiry launches a background goroutine that periodically scans
// KeyTTL for expired keys and removes them from the store.
// It processes keys in chunks of chunkSize to avoid long pauses.
// If the expired ratio within a chunk exceeds expiredRatioThreshold, it
// immediately runs another round rather than waiting for the next tick —
// this handles bursts of expiring keys efficiently.
// The cleanup interval is derived from ServerConfig.Hz and is re-read on
// each tick, so CONFIG SET hz takes effect without a restart.
// The goroutine exits when stop is closed.
func StartActiveExpiry(stop <-chan struct{}) {
	go func() {
		for {
			// A fresh ticker is created on every iteration so that changes to hz
			// via CONFIG SET take effect on the next cycle without a restart.
			// Reusing a single ticker would lock in the interval set at startup.
			interval := ServerConfig.CleanupInterval()
			ticker := time.NewTicker(interval)
			select {
			case <-stop:
				ticker.Stop()
				return
			case <-ticker.C:
				ticker.Stop()
			}

			if !ServerConfig.ActiveExpireEnabled {
				continue
			}
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
