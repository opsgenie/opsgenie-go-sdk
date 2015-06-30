package test

import (
	"testing"
	"time"
)

type WaitAction interface {
	check(t *testing.T) bool
}

func waitFor(t *testing.T, action WaitAction) bool {
	totalSleepMillis := int64(10000)
	currentTimeInMillis := time.Now().UnixNano() / 1000000
	for ((time.Now().UnixNano() / 1000000) - currentTimeInMillis) < totalSleepMillis {
		if action.check(t) {
			return true
		}
		time.Sleep(10 * time.Millisecond)
	}
	return false
}
