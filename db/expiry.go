package db

import (
	"math/rand"
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

// cleanExpiredKeys samples up to chunkSize keys from KeyTTL in random order,
// deletes any that have expired, and returns the number of keys sampled and
// the number deleted.
// Randomisation ensures that all keys get a fair chance of being checked
// regardless of insertion order, avoiding the bias that sync.Map.Range's
// consistent iteration order would otherwise introduce.
func cleanExpiredKeys() (total int, expired int) {
	// Collect all keys that have a TTL set.
	var keys []any
	KeyTTL.Range(func(key, _ any) bool {
		keys = append(keys, key)
		return true
	})

	// Shuffle so every key has an equal chance of landing in the sample window.
	rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })

	now := time.Now().UnixMilli()
	for _, key := range keys {
		if total >= chunkSize {
			break
		}
		total++
		value, ok := KeyTTL.Load(key)
		if ok && value != nil && value.(int64) < now {
			KeyTTL.Delete(key)
			DB.Delete(key)
			expired++
		}
	}
	return
}
