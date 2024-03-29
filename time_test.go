package gofast

import (
	"testing"
	"time"
)

func TestVietnamTime(t *testing.T) {
	now := VietnamTimeNow()
	nowStr := now.Format(time.RFC3339)
	if nowStr[len(nowStr)-6:] != "+07:00" {
		t.Error()
	}

	today := VietnamDateNowIso()
	t.Logf(today)
}

func BenchmarkVietnamTimeLoc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		VietnamDateNowIso()
	}
}
